package utils

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type Logger struct{}

func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	log.Infof(ctx, format, args...)
}

func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(ctx, format, args...)
}
