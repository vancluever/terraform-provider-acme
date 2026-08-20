package main

import (
	"context"
	"log"
	"log/slog"
	"strings"

	legoLog "github.com/go-acme/lego/v5/log"
)

// legoLogHandler is a proxy for lego log messages. This allows messages
// for lego itself to be printed as actual plugin debug logs.
//
// Note that this log proxy only passes through messages, ignores log
// levels, and does not exit on Fatal or its similar functions.  All
// messages are logged at the debug level.
type legoLogHandler struct{}

// initLegoLogger initializes the logger and sets it up as an
// override for the lego standard logger, which is ignored by
// Terraform. This should be run in main, after lego has had a chance
// to initialize the package singleton.
func initLegoLogger() {
	l := &legoLogHandler{}
	legoLog.SetDefault(slog.New(l))
	l.log("Messages from the lego library will show up as DEBUG messages.")
}

func (l *legoLogHandler) Enabled(context.Context, slog.Level) bool { return true }
func (l *legoLogHandler) WithAttrs([]slog.Attr) slog.Handler       { return l }
func (l *legoLogHandler) WithGroup(string) slog.Handler            { return l }
func (l *legoLogHandler) Handle(ctx context.Context, r slog.Record) error {
	l.log(r.Message)
	return nil
}

// log logs the raw message sent to it to the Terraform logger with a
// prefix indicating it came from lego.
//
// All messages are logged at the debug level.
func (l *legoLogHandler) log(args ...any) {
	// Strip any lego-based log level from the string. This should
	// always be in the first argument.
	if len(args) > 0 {
		if _, ok := args[0].(string); ok {
			switch {
			case strings.HasPrefix(args[0].(string), "[INFO] "):
				args[0] = strings.TrimPrefix(args[0].(string), "[INFO] ")

			case strings.HasPrefix(args[0].(string), "[WARN] "):
				args[0] = strings.TrimPrefix(args[0].(string), "[WARN] ")
			}
		}
	}

	log.Println(append([]any{"[DEBUG]", "lego:"}, args...)...)
}
