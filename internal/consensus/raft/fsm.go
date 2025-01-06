package raft

import (
	"encoding/json"
	"io"
	"sync"

	tmlog "github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/types"
	"github.com/hashicorp/raft"
)

type ProposerFSM struct {
	mtx    sync.RWMutex
	state  *ProposerState
	logger tmlog.Logger
}

func NewProposerFSM(logger tmlog.Logger) *ProposerFSM {
	return &ProposerFSM{
		state: &ProposerState{
			Height: 1,
			Round:  0,
		},
		logger: logger,
	}
}

// Apply implements raft.FSM
func (f *ProposerFSM) Apply(log *raft.Log) interface{} {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	var cmd ProposerCommand
	if err := json.Unmarshal(log.Data, &cmd); err != nil {
		f.logger.Error("failed to unmarshal proposer command", "error", err)
		return nil
	}

	switch cmd.Type {
	case "update_proposer":
		f.state.CurrentProposer = cmd.Proposer
		f.state.Height = cmd.Height
		f.state.Round = cmd.Round
	}

	return f.state
}

// Snapshot implements raft.FSM
func (f *ProposerFSM) Snapshot() (raft.FSMSnapshot, error) {
	f.mtx.RLock()
	defer f.mtx.RUnlock()

	return &ProposerSnapshot{
		state: *f.state,
	}, nil
}

// Restore implements raft.FSM
func (f *ProposerFSM) Restore(rc io.ReadCloser) error {
	defer rc.Close()

	var state ProposerState
	if err := json.NewDecoder(rc).Decode(&state); err != nil {
		return err
	}

	f.mtx.Lock()
	f.state = &state
	f.mtx.Unlock()

	return nil
}

type ProposerSnapshot struct {
	state ProposerState
}

func (s *ProposerSnapshot) Persist(sink raft.SnapshotSink) error {
	err := json.NewEncoder(sink).Encode(s.state)
	if err != nil {
		sink.Cancel()
		return err
	}
	return sink.Close()
}

func (s *ProposerSnapshot) Release() {}

type ProposerCommand struct {
	Type     string
	Proposer types.Address
	Height   int64
	Round    int32
}
