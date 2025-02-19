/*
Copyright © 2025 leig <leigme@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	config "github.com/leigme/search/config"
	util "github.com/leigme/search/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg            *config.Json
	path, key, sep string
	clip           bool
)

var rootCmd = &cobra.Command{
	Use:   "search",
	Short: "Configuration file search tool",
	Long: `Configuration file search tool, for example:
	search --p [your config path] --k [the key to be searched for config] --s [separator] --c [true]`,
	Run: func(cmd *cobra.Command, args []string) {
		defer cfg.Update()
		if strings.EqualFold(key, "") {
			log.Fatalln(errors.New("--k can not nil"))
		}
		cfg.Load()
		log.Println(cfg)
		if strings.EqualFold(path, "") {
			if len(cfg.Files) == 0 {
				log.Fatalln(errors.New("--p can not nil"))
			}
			path = cfg.Files[len(cfg.Files)-1].Path
		} else {
			cfg.Add(path)
		}
		keys := strings.Split(key, ",")
		result := search(path, keys...)
		out(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	bingArgs()
	cfg = config.NewJson()
}

func bingArgs() {
	rootCmd.PersistentFlags().StringVar(&path, "p", "", "--p config file path")
	rootCmd.PersistentFlags().StringVar(&key, "k", "", "--k search value by keys")
	rootCmd.PersistentFlags().StringVar(&sep, "s", " ", "--s out content by separator")
	rootCmd.PersistentFlags().BoolVar(&clip, "c", true, "--c is content by clipboard")
}

func search(path string, keys ...string) []string {
	result := make([]string, 0)
	for _, f := range cfg.Files {
		if strings.EqualFold(path, f.Path) {
			values := searchByViper(f, keys...)
			if len(values) > 0 {
				result = append(result, values...)
			}
		}
	}
	return result
}

func searchByViper(file config.File, keys ...string) []string {
	viper.SetConfigType(file.Type)
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	result := make([]string, 0)
	for _, key := range keys {
		for _, k := range viper.AllKeys() {
			if strings.HasSuffix(k, key) {
				result = append(result, viper.GetString(k))
			}
		}
	}
	return util.Unique(result)
}

func out(values []string) {
	result := strings.Join(values, sep)
	fmt.Println(result)
	if clip {
		clipboard.WriteAll(result)
	}
}
