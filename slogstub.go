// SPDX-License-Identifier: GPL-3.0-or-later

// Package slogstub provides test stubs for the [log/slog] package.
package slogstub

import (
	"context"
	"log/slog"

	"github.com/bassosimone/runtimex"
)

// FuncHandler implements [slog.Handler] using configurable functions.
//
// This type is useful for testing code that emits structured log events,
// allowing tests to capture and inspect [slog.Record] values directly
// rather than parsing serialized output.
//
// Each function field must be set before calling the corresponding method.
// Calling a method when its function field is nil will panic.
type FuncHandler struct {
	// EnabledFunc determines if the handler is enabled for the given level.
	EnabledFunc func(ctx context.Context, level slog.Level) bool

	// HandleFunc processes a log record.
	HandleFunc func(ctx context.Context, record slog.Record) error

	// WithAttrsFunc returns a new handler with additional attributes.
	WithAttrsFunc func(attrs []slog.Attr) slog.Handler

	// WithGroupFunc returns a new handler with a group name.
	WithGroupFunc func(name string) slog.Handler
}

var _ slog.Handler = &FuncHandler{}

// Enabled implements [slog.Handler].
func (h *FuncHandler) Enabled(ctx context.Context, level slog.Level) bool {
	runtimex.Assert(h.EnabledFunc != nil)
	return h.EnabledFunc(ctx, level)
}

// Handle implements [slog.Handler].
func (h *FuncHandler) Handle(ctx context.Context, record slog.Record) error {
	runtimex.Assert(h.HandleFunc != nil)
	return h.HandleFunc(ctx, record)
}

// WithAttrs implements [slog.Handler].
func (h *FuncHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	runtimex.Assert(h.WithAttrsFunc != nil)
	return h.WithAttrsFunc(attrs)
}

// WithGroup implements [slog.Handler].
func (h *FuncHandler) WithGroup(name string) slog.Handler {
	runtimex.Assert(h.WithGroupFunc != nil)
	return h.WithGroupFunc(name)
}
