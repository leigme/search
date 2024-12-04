package cmd

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/leigme/search/config"
	"github.com/leigme/search/logger"
	"github.com/leigme/search/model"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set config content",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := model.Config{}
		if bytes, err := os.ReadFile(config.ConfigPath()); err == nil {
			err = json.Unmarshal(bytes, &cfg)
			if err != nil {
				logger.Fatalln(err)
			}
		}
		if !strings.EqualFold(param.Config.Path, "") {
			cfg.Path = param.Config.Path
		}
		if !strings.EqualFold(param.Config.LogPath, "") {
			cfg.LogPath = param.Config.LogPath
		}
		if !strings.EqualFold(param.Config.LogLevel, "") {
			cfg.LogLevel = param.Config.LogLevel
		}
		if bytes, err := json.Marshal(cfg); err == nil {
			err = os.WriteFile(config.ConfigPath(), bytes, os.ModePerm)
			if err != nil {
				logger.Fatalln(err)
			}
		}
	},
}
