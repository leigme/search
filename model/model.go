package model

import (
	"log/slog"
	"strings"
)

var Local Config

type Config struct {
	Path     string `json:"path"`
	LogPath  string `json:"log_path"`
	LogLevel string `json:"log_level"`
}

func (c *Config) GetLogLevel() slog.Level {
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
