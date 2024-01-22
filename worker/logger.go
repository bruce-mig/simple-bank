package worker

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) Print(level zerolog.Level, args ...any) {
	log.WithLevel(level).Msg(fmt.Sprint(args...))
}

func (logger *Logger) Printf(ctx context.Context, format string, v ...any) {
	log.WithLevel(zerolog.DebugLevel).Msgf(format, v...)
}

func (logger *Logger) Debug(args ...any) {
	logger.Print(zerolog.DebugLevel, args...)
}

func (logger *Logger) Info(args ...any) {
	logger.Print(zerolog.InfoLevel, args...)
}

func (logger *Logger) Warn(args ...any) {
	logger.Print(zerolog.WarnLevel, args...)
}

func (logger *Logger) Error(args ...any) {
	logger.Print(zerolog.ErrorLevel, args...)
}

func (logger *Logger) Fatal(args ...any) {
	logger.Print(zerolog.FatalLevel, args...)
}
