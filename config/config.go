package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const configName = "conf.json"

type Json struct {
	Files []File `json:"files"`
}

type File struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

func NewJson() *Json {
	return &Json{}
}

func (j *Json) Load() {
	checkDir(Path())
	data := readFile(Path())
	err := json.Unmarshal(data, j)
	if err != nil {
		log.Fatalln(err)
	}
}

func (j *Json) Update() {
	checkDir(Path())
	data, err := json.Marshal(j)
	if err != nil {
		log.Fatalln(err)
	}
	writeFile(data, Path())
}

func checkDir(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func readFile(path string) []byte {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		log.Fatalln(err)
	}
	stat, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(f).Read(bs)
	return bs
}

func writeFile(data []byte, path string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		log.Fatalln(err)
	}
	w := bufio.NewWriter(f)
	_, err = w.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
	if err = w.Flush(); err != nil {
		log.Fatalln(err)
	}
}

func (j *Json) AddByType(p, t string) {
	if j.Files == nil {
		j.Files = make([]File, 0)
	}
	j.Files = append(j.Files, File{p, t})
}

func (j *Json) Add(path string) {
	ext := filepath.Ext(path)
	if strings.EqualFold(ext, "") {
		log.Fatalln(path, errors.New("The file name has no suffix"))
	}
	j.AddByType(path, ext)
}

func Dir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	executable, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	return filepath.Join(homeDir, ".config", executable)
}

func Path() string {
	return filepath.Join(Dir(), "conf.json")
}
