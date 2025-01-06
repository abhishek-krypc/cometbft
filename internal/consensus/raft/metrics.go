package raft

import (
	"strconv"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/hashicorp/raft"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	raftAppliedIndex    metrics.Gauge
	raftLastLogIndex    metrics.Gauge
	raftLastLogTerm     metrics.Gauge
	raftCommitIndex     metrics.Gauge
	raftPeerCount       metrics.Gauge
	raftLeaderChanges   metrics.Counter
	raftProposerChanges metrics.Counter
}

func NewMetrics() *Metrics {
	const namespace = "cometbft"
	const subsystem = "raft"

	return &Metrics{
		raftAppliedIndex: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "applied_index",
			Help:      "Last applied Raft log index",
		}, []string{}),

		raftLastLogIndex: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "last_log_index",
			Help:      "Last Raft log index",
		}, []string{}),

		raftLastLogTerm: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "last_log_term",
			Help:      "Last Raft log term",
		}, []string{}),

		raftCommitIndex: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "commit_index",
			Help:      "Raft commit index",
		}, []string{}),

		raftPeerCount: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "peer_count",
			Help:      "Number of Raft peers",
		}, []string{}),

		raftLeaderChanges: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "leader_changes_total",
			Help:      "Total number of Raft leader changes",
		}, []string{}),

		raftProposerChanges: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "proposer_changes_total",
			Help:      "Total number of proposer changes",
		}, []string{}),
	}
}

func (m *Metrics) UpdateRaftMetrics(r *raft.Raft) {
	if stats := r.Stats(); stats != nil {
		appliedIndex, _ := strconv.ParseUint(stats["applied_index"], 10, 64)
		lastLogIndex, _ := strconv.ParseUint(stats["last_log_index"], 10, 64)
		lastLogTerm, _ := strconv.ParseUint(stats["last_log_term"], 10, 64)
		commitIndex, _ := strconv.ParseUint(stats["commit_index"], 10, 64)

		m.raftAppliedIndex.Set(float64(appliedIndex))
		m.raftLastLogIndex.Set(float64(lastLogIndex))
		m.raftLastLogTerm.Set(float64(lastLogTerm))
		m.raftCommitIndex.Set(float64(commitIndex))
		// m.raftPeerCount.Set(float64(len(r.Snapshot())))
	}
}
