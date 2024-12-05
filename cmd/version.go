package cmd

import (
	"fmt"
	"os"

	"github.com/leigme/search/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show app version and config content",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if bytes, err := os.ReadFile(config.DefaultConfigPath()); err == nil {
			fmt.Printf(format, "v1.0.2", config.DefaultConfigPath(), string(bytes))
		}
	},
}

const format string = `version: %s\n
config path: %s\n
`
