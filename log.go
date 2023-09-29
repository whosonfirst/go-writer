package writer

import (
	"io"
	"log/slog"
)

// DefaultLogger() returns a `slog.Logger` instance that writes to `io.Discard`.
func DefaultLogger() *slog.Logger {
	return slog.New(io.Discard, "", log.Lshortfile)
}
