package raft

import (
	"fmt"
	"sync"
	"time"

	"github.com/cometbft/cometbft/config"
	tmlog "github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/types"
	"github.com/hashicorp/raft"
)

type RaftProposer struct {
	mtx         sync.RWMutex
	raft        *raft.Raft
	config      *config.ConsensusConfig
	nodeInfo    p2p.NodeInfo
	raftConfig  *raft.Config
	store       *RaftStore
	fsm         *ProposerFSM
	transport   *raft.NetworkTransport
	metrics     *Metrics
	logger      tmlog.Logger
	raftLogger  *RaftLogger
	initialized bool
}

// Encode encodes the ProposerCommand into a byte slice
func (cmd *ProposerCommand) Encode() []byte {
	return []byte(fmt.Sprintf("%s:%d:%d", cmd.Type, cmd.Height, cmd.Round))
}

// Proposer state to be stored in Raft
type ProposerState struct {
	CurrentProposer types.Address
	Height          int64
	Round           int32
}

// NewRaftProposer creates a new Raft proposer instance
func NewRaftProposer(
	cfg *config.ConsensusConfig,
	logger tmlog.Logger,
	store *RaftStore,
	metrics *Metrics,
) (*RaftProposer, error) {
	if logger == nil {
		logger = tmlog.NewNopLogger()
	}

	// Create Raft-compatible logger
	raftLogger := NewRaftLogger(logger)

	rp := &RaftProposer{
		config:     cfg,
		store:      store,
		metrics:    metrics,
		logger:     logger,
		raftLogger: raftLogger,
	}

	// Create FSM with CometBFT logger
	rp.fsm = NewProposerFSM(logger)

	// Initialize Raft configuration with Raft-compatible logger
	raftConfig := raft.DefaultConfig()
	raftConfig.Logger = raftLogger
	raftConfig.SnapshotInterval = 24 * time.Hour
	raftConfig.SnapshotThreshold = 2048
	raftConfig.TrailingLogs = 10240

	if cfg.RaftConfig.EnableMetrics {
		// raftConfig.MetricsRegistry = metrics.registry
	}

	rp.raftConfig = raftConfig

	return rp, nil
}

// UpdateNodeInfo updates the node info and completes Raft initialization
func (rp *RaftProposer) UpdateNodeInfo(nodeInfo p2p.NodeInfo, bootstrap bool, joinAddress string) error {
	rp.mtx.Lock()
	defer rp.mtx.Unlock()

	rp.nodeInfo = nodeInfo
	rp.raftConfig.LocalID = raft.ServerID(nodeInfo.ID())

	addr := rp.config.RaftConfig.ListenAddress

	// Create transport with Raft-compatible logger
	transport, err := raft.NewTCPTransport(
		addr,
		nil,
		3,
		10*time.Second,
		rp.raftLogger,
	)
	if err != nil {
		return fmt.Errorf("failed to create Raft transport: %w", err)
	}
	rp.transport = transport

	// Create Raft instance
	r, err := raft.NewRaft(
		rp.raftConfig,
		rp.fsm,
		rp.store.logStore,
		rp.store.stableStore,
		rp.store.snapStore,
		rp.transport,
	)
	if err != nil {
		return fmt.Errorf("failed to create Raft instance: %w", err)
	}
	rp.raft = r

	rp.logger.Info("Initialized Raft proposer",
		"nodeID", nodeInfo.ID(),
		"address", addr,
	)

	rp.initialized = true

	// Bootstrap or join cluster based on configuration
	if bootstrap {
		return rp.Bootstrap()
	} else if joinAddress != "" {
		return rp.JoinCluster(joinAddress)
	}
	return nil
}

// GetCurrentProposer returns the current proposer based on Raft consensus
func (rp *RaftProposer) GetCurrentProposer(height int64, round int32) (types.Address, error) {
	if !rp.initialized {
		return nil, fmt.Errorf("raft proposer not initialized - call UpdateNodeInfo first")
	}

	rp.mtx.RLock()
	defer rp.mtx.RUnlock()

	// if rp.raft.State() != raft.Leader {
	// 	leader := rp.raft.Leader()
	// 	if leader == "" {
	// 		return nil, fmt.Errorf("no leader available")
	// 	}
	// 	return nil, fmt.Errorf("not leader - forward to %s", leader)
	// }

	cmd := ProposerCommand{
		Type:   "get_proposer",
		Height: height,
		Round:  round,
	}

	future := rp.raft.Apply(cmd.Encode(), 5*time.Second)
	if err := future.Error(); err != nil {
		rp.logger.Error("Failed to apply proposer command",
			"height", height,
			"round", round,
			"error", err,
		)
		return nil, fmt.Errorf("failed to apply proposer command: %w", err)
	}

	if resp, ok := future.Response().(*ProposerState); ok {
		return resp.CurrentProposer, nil
	}
	return nil, fmt.Errorf("invalid response type from FSM")
}

// Bootstrap starts a new Raft cluster if this is the first node
func (rp *RaftProposer) Bootstrap() error {
	// fmt.Println("print infor node id : " + rp.nodeInfo)
	exist, err := rp.store.CheckRaftStoreExists(&rp.config.RaftConfig)
	if err != nil {
		return err
	} else if exist {
		return nil
	}
	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raft.ServerID(rp.nodeInfo.ID()),
				Address: rp.transport.LocalAddr(),
			},
		},
	}

	return rp.raft.BootstrapCluster(configuration).Error()
}

// JoinCluster adds this node to an existing Raft cluster
func (rp *RaftProposer) JoinCluster(leaderAddr string) error {
	if !rp.initialized {
		return fmt.Errorf("raft proposer not initialized - call UpdateNodeInfo first")
	}

	// Connect to leader
	// conn, err := rp.transport.Connect(raft.ServerAddress(leaderAddr), 10*time.Second)
	// if err != nil {
	//     return fmt.Errorf("failed to connect to leader: %w", err)
	// }
	// defer conn.Close()

	// Send join request
	future := rp.raft.AddVoter(
		raft.ServerID(rp.nodeInfo.ID()),
		raft.ServerAddress(rp.transport.LocalAddr()),
		0,
		0,
	)
	return future.Error()
}

// // GetCurrentProposer returns the current proposer based on Raft consensus
// func (rp *RaftProposer) GetCurrentProposer(height int64, round int32) (types.ValidatorAddress, error) {
// 	rp.mtx.RLock()
// 	defer rp.mtx.RUnlock()

// 	// Get the current state from Raft
// 	state := &ProposerState{}
// 	if err := rp.raft.Apply(state, 10*time.Second).Error(); err != nil {
// 		return nil, fmt.Errorf("failed to get proposer from Raft: %v", err)
// 	}

// 	return state.CurrentProposer, nil
// }

func (rp *RaftProposer) GetLogger() tmlog.Logger {
	return rp.logger
}
