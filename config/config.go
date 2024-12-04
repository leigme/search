package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

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

func DefaultConfigPath() string {
	return filepath.Join(homeDir(), workDir, configName)
}

func InitConfig() {
	if err := os.RemoveAll(DefaultConfigPath()); err != nil {
		log.Println(err)
	}
	err := os.MkdirAll(filepath.Join(homeDir(), workDir), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = os.Create(DefaultConfigPath())
	if err != nil {
		log.Fatalln(err)
	}
	cfg := model.Config{
		Path: DefaultConfigPath(),
	}
	SaveConfig(&cfg)
}

func LoadConfig(cfg *model.Config) {
	_, err := os.Stat(DefaultConfigPath())
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		InitConfig()
	}
	data, err := os.ReadFile(DefaultConfigPath())
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		log.Fatalln(err)
	}
}

func SaveConfig(cfg *model.Config) {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(DefaultConfigPath(), bytes, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}
