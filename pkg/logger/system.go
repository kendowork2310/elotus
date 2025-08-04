package logger

import "go.uber.org/zap"

var l *zap.Logger

func System() *zap.Logger {
	if l == nil {
		l, _ = zap.NewProduction()
		l.Named("SYSTEM")
		return l
	} else {
		return l
	}
}
