package raft

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cometbft/cometbft/config"
	cmtos "github.com/cometbft/cometbft/internal/os"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

type RaftStore struct {
	logStore    *raftboltdb.BoltStore
	stableStore *raftboltdb.BoltStore
	snapStore   raft.SnapshotStore
	baseDir     string
}

func NewRaftStore(cfg *config.RaftConfig) (*RaftStore, error) {
	baseDir := filepath.Join(cfg.RaftDir, "raft")

	err := cmtos.EnsureDir(filepath.Dir(baseDir), 0o700)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure WAL directory is in place: %w", err)
	}
	// Create BoltDB-backed log store
	logPath := filepath.Join(baseDir, "raft-log.db")
	logStore, err := raftboltdb.NewBoltStore(logPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create raft log store: %v", err)
	}

	// Create BoltDB-backed stable store
	stablePath := filepath.Join(baseDir, "raft-stable.db")
	stableStore, err := raftboltdb.NewBoltStore(stablePath)
	if err != nil {
		logStore.Close()
		return nil, fmt.Errorf("failed to create raft stable store: %v", err)
	}

	// Create file-based snapshot store
	snapPath := filepath.Join(baseDir, "snapshots")
	snapStore, err := raft.NewFileSnapshotStore(snapPath, 3, nil)
	if err != nil {
		logStore.Close()
		stableStore.Close()
		return nil, fmt.Errorf("failed to create snapshot store: %v", err)
	}

	return &RaftStore{
		logStore:    logStore,
		stableStore: stableStore,
		snapStore:   snapStore,
		baseDir:     baseDir,
	}, nil
}

func (rs *RaftStore) Close() error {
	if err := rs.logStore.Close(); err != nil {
		return err
	}
	return rs.stableStore.Close()
}

// CheckRaftStoreExists checks if the Raft store exists and is properly set up.
func (rs *RaftStore) CheckRaftStoreExists(cfg *config.RaftConfig) (bool, error) {
	baseDir := filepath.Join(cfg.RaftDir, "raft")

	// Check if the base directory exists
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		return false, nil // Base directory does not exist
	}

	// Check if the Raft log store exists
	logPath := filepath.Join(baseDir, "raft-log.db")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		return false, nil // Log store does not exist
	}

	// Check if the Raft stable store exists
	stablePath := filepath.Join(baseDir, "raft-stable.db")
	if _, err := os.Stat(stablePath); os.IsNotExist(err) {
		return false, nil // Stable store does not exist
	}

	// Check if the snapshot store directory exists
	snapPath := filepath.Join(baseDir, "snapshots")
	if _, err := os.Stat(snapPath); os.IsNotExist(err) {
		return false, nil // Snapshot store does not exist
	}

	// All required components exist
	return true, nil
}
