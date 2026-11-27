// Package gatewaylog adapts Beego's logs package to the shared slog JSON handler so
// that existing `logs.*` calls emit structured JSON through slog with no call-site
// changes, keeping a single log format on stdout.
package gatewaylog

import (
	"context"
	"log/slog"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// SlogJSONAdapter implements logs.Logger by forwarding every message to the slog
// default logger. It is registered and enabled in main.go.
type SlogJSONAdapter struct{}

func (a *SlogJSONAdapter) Init(config string) error { return nil }
func (a *SlogJSONAdapter) Destroy()                  {}
func (a *SlogJSONAdapter) Flush()                    {}

// WriteMsg receives the already-formatted message from Beego and re-emits it via
// slog. Beego performs its own level filtering, so we map to slog levels; the slog
// handler performs the final level filtering based on its configured level.
func (a *SlogJSONAdapter) WriteMsg(when time.Time, msg string, level int) error {
	slog.Log(context.Background(), mapLevel(level), msg,
		slog.String("service", "gateway"))
	return nil
}

func mapLevel(level int) slog.Level {
	switch level {
	case logs.LevelEmergency, logs.LevelAlert, logs.LevelCritical, logs.LevelError:
		return slog.LevelError
	case logs.LevelWarning:
		return slog.LevelWarn
	case logs.LevelNotice, logs.LevelInfo:
		return slog.LevelInfo
	case logs.LevelDebug, logs.LevelTrace:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

// Ensure the adapter satisfies the Beego logs.Logger interface.
var _ logs.Logger = (*SlogJSONAdapter)(nil)
