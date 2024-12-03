package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/leigme/search/model"
	homedir "github.com/mitchellh/go-homedir"
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

func ConfigPath() string {
	return filepath.Join(homeDir(), workDir, configName)
}

func InitConfig() {
	err := os.MkdirAll(filepath.Join(homeDir(), workDir), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = os.Create(ConfigPath())
	if err != nil {
		log.Fatalln(err)
	}
	SaveConfig()
}

func LoadConfig(cfg *model.Config) {
	_, err := os.Stat(ConfigPath())
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		InitConfig()
	}
	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &model.Local)
	if err != nil {
		log.Fatalln(err)
	}
	if !strings.EqualFold(cfg.Path, "") {
		model.Local.Path = cfg.Path
	}
	if !strings.EqualFold(cfg.LogPath, "") {
		model.Local.LogPath = cfg.LogPath
	}
	if !strings.EqualFold(cfg.LogLevel, "") {
		model.Local.LogLevel = cfg.LogLevel
	}
}

func SaveConfig() {
	bytes, err := json.Marshal(model.Local)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(ConfigPath(), bytes, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}
