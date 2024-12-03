/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	config "github.com/leigme/search/config"
	logger "github.com/leigme/search/logger"
	model "github.com/leigme/search/model"
	util "github.com/leigme/search/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.design/x/clipboard"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := search(args...); err == nil {
			out(result)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	cfg := model.Config{}
	bingArgs(&cfg)
	config.LoadConfig(&cfg)
	logger.InitLogger(&model.Local)
	rootCmd.AddCommand(initCmd, configCmd)
	config.SaveConfig()
}

func bingArgs(cfg *model.Config) {
	rootCmd.PersistentFlags().StringVar(&cfg.Path, "path", "", "config file path (default is $HOME/.search/config.json)")
	rootCmd.PersistentFlags().StringVar(&cfg.LogPath, "log", "", "config file path (default is $HOME/.search/config.json)")
	rootCmd.PersistentFlags().StringVar(&cfg.LogLevel, "level", "", "")
}

func search(keys ...string) ([]string, error) {
	result := make([]string, 0)
	suffix := filepath.Base(config.ConfigPath())
	var configType string
	switch suffix {
	case "json":
		configType = "json"
	case "ini":
		configType = "ini"
	case "yaml":
		configType = "yaml"
	case "yml":
		configType = "yml"
	default:
		return nil, fmt.Errorf("config: [%s] format nonsupport", suffix)
	}
	viper.SetConfigType(configType)
	viper.ReadInConfig()
	for _, key := range keys {
		for _, k := range viper.AllKeys() {
			if strings.HasSuffix(k, key) {
				result = append(result, viper.GetString(k))
			}
		}
	}
	return util.Unique(result), nil
}

func out(a ...any) {
	content := fmt.Sprintln(a...)
	clipboard.Write(clipboard.FmtText, []byte(content))
	fmt.Println(content)
}

func outf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	clipboard.Write(clipboard.FmtText, []byte(content))
	fmt.Println(content)
}
