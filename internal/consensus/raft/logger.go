package raft

import (
	"io"
	"log"

	tmlog "github.com/cometbft/cometbft/libs/log"
	hclog "github.com/hashicorp/go-hclog"
)

// RaftLogger adapts CometBFT's logger to Raft's logger interface
type RaftLogger struct {
	logger tmlog.Logger
}

// GetLevel implements hclog.Logger.
func (r *RaftLogger) GetLevel() hclog.Level {
	panic("unimplemented")
}

// ImpliedArgs implements hclog.Logger.
func (r *RaftLogger) ImpliedArgs() []interface{} {
	panic("unimplemented")
}

// IsDebug implements hclog.Logger.
func (r *RaftLogger) IsDebug() bool {
	panic("unimplemented")
}

// IsError implements hclog.Logger.
func (r *RaftLogger) IsError() bool {
	panic("unimplemented")
}

// IsInfo implements hclog.Logger.
func (r *RaftLogger) IsInfo() bool {
	panic("unimplemented")
}

// IsTrace implements hclog.Logger.
func (r *RaftLogger) IsTrace() bool {
	panic("unimplemented")
}

// IsWarn implements hclog.Logger.
func (r *RaftLogger) IsWarn() bool {
	panic("unimplemented")
}

// Log implements hclog.Logger.
func (r *RaftLogger) Log(level hclog.Level, msg string, args ...interface{}) {
	panic("unimplemented")
}

// Named implements hclog.Logger.
func (r *RaftLogger) Named(name string) hclog.Logger {
	panic("unimplemented")
}

// ResetNamed implements hclog.Logger.
func (r *RaftLogger) ResetNamed(name string) hclog.Logger {
	panic("unimplemented")
}

// SetLevel implements hclog.Logger.
func (r *RaftLogger) SetLevel(level hclog.Level) {
	panic("unimplemented")
}

// StandardLogger implements hclog.Logger.
func (r *RaftLogger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	panic("unimplemented")
}

// StandardWriter implements hclog.Logger.
func (r *RaftLogger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	panic("unimplemented")
}

// Trace implements hclog.Logger.
func (r *RaftLogger) Trace(msg string, args ...interface{}) {
	panic("unimplemented")
}

// With implements hclog.Logger.
func (r *RaftLogger) With(args ...interface{}) hclog.Logger {
	panic("unimplemented")
}

func NewRaftLogger(logger tmlog.Logger) *RaftLogger {
	return &RaftLogger{
		logger: logger,
	}
}

// Implement raft.Logger interface
func (r *RaftLogger) Printf(format string, args ...interface{}) {
	r.logger.Info(format, args...)
}

func (r *RaftLogger) Println(v ...interface{}) {
	r.logger.Info("msg", v...)
}

// Debug logs are noops - implement if needed
func (r *RaftLogger) Debug(format string, args ...interface{}) {
	r.logger.Debug(format, args...)
}

func (r *RaftLogger) Error(format string, args ...interface{}) {
	r.logger.Error(format, args...)
}

func (r *RaftLogger) Info(format string, args ...interface{}) {
	r.logger.Info(format, args...)
}

func (r *RaftLogger) Warn(format string, args ...interface{}) {
	r.logger.Info(format, args...)
}

// Name returns the name of the logger
func (r *RaftLogger) Name() string {
	return "raft"
}

// SetLogLevel is a no-op as we use CometBFT's log levels
func (r *RaftLogger) SetLogLevel(level string) {}

// GetLogLevel returns the current log level
func (r *RaftLogger) GetLogLevel() string {
	return "info"
}

// Write implements io.Writer
func (r *RaftLogger) Write(p []byte) (n int, err error) {
	r.logger.Info(string(p))
	return len(p), nil
}
