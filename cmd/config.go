package cmd

import (
	"encoding/json"
	"log/slog"

	model "github.com/leigme/search/model"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "cfg",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if bytes, err := json.Marshal(model.Local); err == nil {
			slog.Info(string(bytes))
		}
	},
}
