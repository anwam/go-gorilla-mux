package main

import (
	"os"
	"time"

	"log/slog"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			a.Key = "timestamp"
			a.Value = slog.TimeValue(time.Now())
		}
		return a
	}}))
