package cmd

import "github.com/spf13/cobra"

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := Search(args...); err == nil {
			Out(result)
		}
	},
}
