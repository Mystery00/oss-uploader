package logger

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log/slog"
)

type CliHandler struct {
	slog.Handler
	w           io.Writer
	EnableColor bool
	TimeFormat  string
	LevelText   []string
}

type CliOptions struct {
	Level       slog.Level
	EnableColor bool
	TimeFormat  string
	LevelText   []string
}

func NewCliHandler(w io.Writer, opts *CliOptions) *CliHandler {
	if opts == nil {
		opts = &CliOptions{}
	}
	levelText := []string{"调试", "信息", "警告", "错误"}
	if len(opts.LevelText) == 4 {
		levelText = opts.LevelText
	}
	return &CliHandler{
		Handler:     slog.NewTextHandler(nil, &slog.HandlerOptions{Level: opts.Level}),
		w:           w,
		EnableColor: opts.EnableColor,
		TimeFormat:  opts.TimeFormat,
		LevelText:   levelText,
	}
}

func (h *CliHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *CliHandler) Handle(ctx context.Context, r slog.Record) error {
	msg := make([]any, 0)
	level := h.formatLevel(r.Level)
	msg = append(msg, level)
	if h.TimeFormat != "" {
		msg = append(msg, r.Time.Format(h.TimeFormat))
	}
	msg = append(msg, r.Message)
	r.Attrs(func(attr slog.Attr) bool {
		msg = append(msg, fmt.Sprintf("%s=%v", attr.Key, attr.Value.Any()))
		return true
	})

	_, err := fmt.Fprintln(h.w, msg...)
	return err
}

func (h *CliHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CliHandler{
		Handler:     h.Handler.WithAttrs(attrs),
		EnableColor: h.EnableColor,
		TimeFormat:  h.TimeFormat,
	}
}

func (h *CliHandler) WithGroup(name string) slog.Handler {
	return &CliHandler{
		Handler:     h.Handler.WithGroup(name),
		EnableColor: h.EnableColor,
		TimeFormat:  h.TimeFormat,
	}
}

func (h *CliHandler) formatLevel(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return levelToText(h.EnableColor, color.FgCyan, h.LevelText[0])
	case slog.LevelInfo:
		return levelToText(h.EnableColor, color.FgGreen, h.LevelText[1])
	case slog.LevelWarn:
		return levelToText(h.EnableColor, color.FgYellow, h.LevelText[2])
	case slog.LevelError:
		return levelToText(h.EnableColor, color.FgRed, h.LevelText[3])
	default:
		panic(fmt.Sprintf("unknown level: %v", level))
	}
}

func levelToText(enableColor bool, c color.Attribute, text string) string {
	if enableColor {
		return color.New(color.Bold, c).Sprintf("[%s]", text)
	} else {
		return fmt.Sprintf("[%s]", text)
	}
}
