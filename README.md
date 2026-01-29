# Golang Helpers for slog Testing

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/slogstub)](https://pkg.go.dev/github.com/bassosimone/slogstub) [![Build Status](https://github.com/bassosimone/slogstub/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/slogstub/actions) [![codecov](https://codecov.io/gh/bassosimone/slogstub/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/slogstub)

The `slogstub` Go package contains small helpers for testing code that uses `log/slog`.

For example:

```Go
import (
	"context"
	"log/slog"

	"github.com/bassosimone/slogstub"
	"github.com/stretchr/testify/assert"
)

// Create a handler that captures log records.
var captured []slog.Record
handler := &slogstub.FuncHandler{
	EnabledFunc: func(ctx context.Context, level slog.Level) bool {
		return true
	},
	HandleFunc: func(ctx context.Context, record slog.Record) error {
		captured = append(captured, record)
		return nil
	},
	WithAttrsFunc: func(attrs []slog.Attr) slog.Handler {
		return handler
	},
	WithGroupFunc: func(name string) slog.Handler {
		return handler
	},
}

// Use the handler in code under test.
logger := slog.New(handler)
logger.Info("test message", "key", "value")

// Verify the captured records.
assert.Len(t, captured, 1)
assert.Equal(t, "test message", captured[0].Message)
```

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/slogstub
```

## Development

To run the tests:
```sh
go test -v .
```

To measure test coverage:
```sh
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```
