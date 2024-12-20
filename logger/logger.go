package logger

import (
	"go-cli-template/env"
	"log/slog"
	"os"
)

func InitLog() {
	level := slog.LevelInfo
	if env.Debug {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(NewCliHandler(os.Stderr, &CliOptions{
		Level:       level,
		EnableColor: true,
		LevelText:   []string{"调试", "信息", "警告", "错误"},
	})))
}
