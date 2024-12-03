package cmd

import (
	"os"

	config "github.com/leigme/search/config"
	logger "github.com/leigme/search/logger"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		os.Remove(config.ConfigPath())
		config.InitConfig()
		bytes, err := os.ReadFile(config.ConfigPath())
		if err != nil {
			logger.Fatalln(err)
		}
		slog.Info(string(bytes))
	},
}
