package raft

import (
	"fmt"
	"time"

	"github.com/hashicorp/raft"
)

type CleanupHandler struct {
	raft    *raft.Raft
	store   *RaftStore
	metrics *Metrics
	baseDir string
}

func NewCleanupHandler(r *raft.Raft, store *RaftStore, metrics *Metrics) *CleanupHandler {
	return &CleanupHandler{
		raft:    r,
		store:   store,
		metrics: metrics,
		baseDir: store.baseDir,
	}
}

func (h *CleanupHandler) Cleanup() error {
	// Shutdown Raft gracefully
	if h.raft != nil {
		future := h.raft.Shutdown()
		if err := future.Error(); err != nil {
			return fmt.Errorf("failed to shutdown raft: %v", err)
		}
	}

	// Close stores
	if h.store != nil {
		if err := h.store.Close(); err != nil {
			return fmt.Errorf("failed to close raft stores: %v", err)
		}
	}

	// Remove old log files
	if err := h.cleanupOldLogs(); err != nil {
		return fmt.Errorf("failed to cleanup old logs: %v", err)
	}

	return nil
}

func (h *CleanupHandler) cleanupOldLogs() error {
	// Get list of log files
	// pattern := filepath.Join(h.baseDir, "raft-*.db")
	// matches, err := filepath.Glob(pattern)
	// if err != nil {
	// 	return err
	// }

	// Keep only the most recent files based on modification time
	// if len(matches) > h.store.config.MaxLogFiles {
	//     for _, file := range matches[:len(matches)-h.store.config.MaxLogFiles] {
	//         if err := os.Remove(file); err != nil {
	//             return err
	//         }
	//     }
	// }

	return nil
}

func (h *CleanupHandler) StartPeriodicCleanup(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			if err := h.cleanupOldLogs(); err != nil {
				// h.logger.Error("failed to cleanup old logs", "error", err)
			}
		}
	}()
}
