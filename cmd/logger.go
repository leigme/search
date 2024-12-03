package cmd

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"golang.design/x/clipboard"
)

func initLogger(cfg Config) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       cfg.LogLevel,
		ReplaceAttr: nil,
	}))
	slog.SetDefault(l)
	err := clipboard.Init()
	if err != nil {
		log.Fatalln(err)
	}
}

func Out(a ...any) {
	content := fmt.Sprintln(a...)
	clipboard.Write(clipboard.FmtText, []byte(content))
	fmt.Println(content)
}

func Outf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	clipboard.Write(clipboard.FmtText, []byte(content))
	fmt.Println(content)
}

func Fatalln(err error) {
	slog.Error(err.Error(), err)
	os.Exit(1)
}
