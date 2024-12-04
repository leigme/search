package cmd

import (
	config "github.com/leigme/search/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init config file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
}
