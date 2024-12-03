package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		os.Remove(configPath())
		initConfig()
		bytes, err := os.ReadFile(configPath())
		if err != nil {
			Fatalln(err)
		}
		Out(string(bytes))
	},
}
