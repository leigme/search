package logger

import (
	"log"
	"log/slog"
	"os"

	model "github.com/leigme/search/model"

	"golang.design/x/clipboard"
)

func InitLogger(cfg *model.Config) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       cfg.GetLogLevel(),
		ReplaceAttr: nil,
	}))
	slog.SetDefault(l)
	err := clipboard.Init()
	if err != nil {
		log.Fatalln(err)
	}
}

func Fatalln(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
