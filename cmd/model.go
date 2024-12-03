package cmd

import "log/slog"

type Config struct {
	LogLevel slog.Level `json:"log_level"`
	LogPath  string     `json:"log_path"`
	Path     string     `json:"path"`
}
