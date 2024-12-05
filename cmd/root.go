/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"golang.design/x/clipboard"
	"log"
	"path/filepath"
	"strings"

	config "github.com/leigme/search/config"
	model "github.com/leigme/search/model"
	util "github.com/leigme/search/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var param model.Param
var cfg model.Config

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
		if strings.EqualFold(param.Keys, "") {
			fmt.Println("search key is nil")
			return
		}
		if !strings.EqualFold(param.File, "") {
			cfg.Path = param.File
		}
		keys := strings.Split(param.Keys, ",")
		if result, err := search(keys...); err == nil {
			out(result)
		}
		config.SaveConfig(&cfg)
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
	param = model.Param{}
	bingArgs(&param)
	cfg = model.Config{}
	config.LoadConfig(&cfg)
	if util.IsLinux() {
		if err := clipboard.Init(); err != nil {
			log.Fatalln(err)
		}
	}
	rootCmd.AddCommand(initCmd, setCmd, versionCmd)
}

func bingArgs(param *model.Param) {
	rootCmd.Flags().StringVar(&param.Keys, "keys", "", "")
	rootCmd.Flags().StringVar(&param.File, "file", "", "search file absolute path")
	rootCmd.Flags().StringVar(&param.Clip, "clip", "", "search clipboard content")
	setCmd.Flags().StringVar(&param.Config.ValueSplit, "vs", "", "value split")
	setCmd.Flags().StringVar(&param.Config.Path, "path", "", "config path (default path: $HOME/.search/config.json)")
	setCmd.Flags().StringVar(&param.Config.LogPath, "log", "", "")
	setCmd.Flags().StringVar(&param.Config.LogLevel, "level", "", "")
}

func search(keys ...string) ([]string, error) {
	result := make([]string, 0)
	suffix := filepath.Ext(cfg.Path)
	var configType string
	switch suffix {
	case ".json":
		configType = "json"
	case ".ini":
		configType = "ini"
	case ".yaml":
		configType = "yaml"
	case ".yml":
		configType = "yaml"
	default:
		return nil, fmt.Errorf("config: [%s] format nonsupport", suffix)
	}
	viper.SetConfigType(configType)
	viper.SetConfigFile(cfg.Path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	for _, key := range keys {
		for _, k := range viper.AllKeys() {
			if strings.HasSuffix(k, key) {
				if strings.EqualFold(cfg.ValueSplit, "") {
					result = append(result, viper.GetString(k))
				} else {
					values := strings.Split(viper.GetString(k), cfg.ValueSplit)
					result = append(result, values...)
				}
			}
		}
	}
	return util.Unique(result), nil
}

func out(a ...any) {
	content := fmt.Sprintln(a...)
	if util.IsLinux() {
		clipboard.Write(clipboard.FmtText, []byte(content))
	}
	fmt.Println(content)
}

func outf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	if util.IsLinux() {
		clipboard.Write(clipboard.FmtText, []byte(content))
	}
	fmt.Println(content)
}
