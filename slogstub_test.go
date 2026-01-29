//
// SPDX-License-Identifier: GPL-3.0-or-later
//

package slogstub

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncHandler(t *testing.T) {
	wantErr := errors.New("mocked error")
	wantHandler := &FuncHandler{}

	handler := &FuncHandler{
		EnabledFunc: func(ctx context.Context, level slog.Level) bool {
			return level >= slog.LevelWarn
		},
		HandleFunc: func(ctx context.Context, record slog.Record) error {
			return wantErr
		},
		WithAttrsFunc: func(attrs []slog.Attr) slog.Handler {
			return wantHandler
		},
		WithGroupFunc: func(name string) slog.Handler {
			return wantHandler
		},
	}

	ctx := context.Background()

	assert.False(t, handler.Enabled(ctx, slog.LevelInfo))
	assert.True(t, handler.Enabled(ctx, slog.LevelWarn))
	assert.Equal(t, wantErr, handler.Handle(ctx, slog.Record{}))
	assert.Same(t, wantHandler, handler.WithAttrs(nil))
	assert.Same(t, wantHandler, handler.WithGroup("group"))
}
