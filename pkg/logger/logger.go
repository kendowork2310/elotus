package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	ErrKey        = "_err"
	InfoKey       = "_info"
	DebugKey      = "_debug"
	WarningKey    = "_warning"
	CustomDataKey = "_custom_data"
)

const (
	DebugLevel   = "DEBUG"
	InfoLevel    = "INFO"
	WarningLevel = "WARNING"
	ErrorLevel   = "ERROR"
)

type LogArr []string

func Gin(c *gin.Context) Logger {
	return &logger{
		GinCtx: c,
	}
}

type Logger interface {
	With(key string, value interface{}) Logger
	Error(format string, args ...interface{})
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Warning(format string, args ...interface{})
	RequestInfo(infoType, value string, args ...interface{})
}

type logger struct {
	GinCtx *gin.Context
}

func (l *logger) With(key string, value interface{}) Logger {
	if value != nil {
		l.Append(CustomDataKey, key)
		l.Set(key, value)
	}
	return l
}

func (l *logger) Error(format string, args ...interface{}) {
	l.Append(ErrKey, format, args...)
}

func (l *logger) Debug(format string, args ...interface{}) {
	l.Append(DebugKey, format, args...)
}

func (l *logger) Info(format string, args ...interface{}) {
	l.Append(InfoKey, format, args...)
}

func (l *logger) Warning(format string, args ...interface{}) {
	l.Append(WarningKey, format, args...)
}

func (l *logger) RequestInfo(infoType, value string, args ...interface{}) {
	l.Append(infoType, value, args...)
}

func (l *logger) Set(key string, value interface{}) {
	if l.GinCtx != nil {
		l.GinCtx.Set(key, value)
	}
}

func (l *logger) Get(key string) interface{} {
	if l.GinCtx != nil {
		value, _ := l.GinCtx.Get(key)
		return value
	}
	return nil
}

func (l *logger) GetLogData(key string) LogArr {
	val := l.Get(key)
	if val != nil {
		return val.(LogArr)
	}
	return nil
}

func (l *logger) Initial(key string) {
	if l.GinCtx != nil {
		val, exists := l.GinCtx.Get(key)
		if !exists || val == nil {
			l.GinCtx.Set(key, LogArr{})
		}
	}
}

func (l *logger) Append(key string, format string, args ...interface{}) {
	l.Initial(key)
	if value := l.Get(key); value != nil {
		logArr := value.(LogArr)
		logArr = append(logArr, fmt.Sprintf(format, args...))
		l.Set(key, logArr)
	}
}
