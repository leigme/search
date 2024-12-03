package cmd

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const workDir = ".search"

const configName = "config.json"

func homeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}
	return home
}

func configPath() string {
	return filepath.Join(homeDir(), workDir, configName)
}

func initConfig() {
	err := os.MkdirAll(filepath.Join(homeDir(), workDir), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = os.Create(configPath())
	if err != nil {
		log.Fatalln(err)
	}
	saveConfig()
}

func loadConfig() {
	_, err := os.Stat(configPath())
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		initConfig()
	}
	data, err := os.ReadFile(configPath())
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
}

func saveConfig() {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(configPath(), bytes, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}

func Search(keys ...string) ([]string, error) {
	result := make([]string, 0)
	suffix := filepath.Base(configPath())
	viper.SetConfigName(configPath())
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
		return nil, errors.New(fmt.Sprintf("config: [%s] format nonsupport", suffix))
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
	return Unique[string](result), nil
}

func Unique[T cmp.Ordered](ss []T) []T {
	size := len(ss)
	if size == 0 {
		return []T{}
	}
	newSlices := make([]T, 0)
	m1 := make(map[T]byte)
	for _, v := range ss {
		if _, ok := m1[v]; !ok {
			m1[v] = 1
			newSlices = append(newSlices, v)
		}
	}
	return newSlices
}
