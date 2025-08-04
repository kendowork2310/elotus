package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	HeaderXRequestID     = "X-Request-ID"
	HeaderCFConnectingIP = "CF-Connecting-IP"
	HeaderXForwardedFor  = "X-Forwarded-For"
	HeaderTrueClientIP   = "True-Client-IP"
)

const (
	XRequestID      = "x_request_id"
	CFConnectingIP  = "cf_connecting_ip"
	XForwardedFor   = "x_forwarded_for"
	TrueClientIP    = "true_client_ip"
	ContextResponse = "x-response"
)

func init() {
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "_datetime"
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.MessageFieldName = "message"
}

func RequestInfo(appName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(HeaderXRequestID)
		if reqID == "" {
			reqID = uuid.New().String()
		}

		c.Set(string(contextKeyRequestID), reqID)

		var (
			defaultLogLvl  = zerolog.InfoLevel
			startTime      = time.Now()
			reqEndpoint    = c.Request.URL.Path
			reqQueries     = c.Request.URL.RawQuery
			reqMethod      = c.Request.Method
			handlerName    = c.HandlerName()
			xRequestId     = reqID
			xForwardedFor  = c.GetHeader(HeaderXForwardedFor)
			trueClientIP   = c.GetHeader(HeaderTrueClientIP)
			cfConnectingIP = c.GetHeader(HeaderCFConnectingIP)
		)

		zrLogger := log.With().
			Str("app_name", appName).
			Str("endpoint", reqEndpoint).
			Str("method", reqMethod).
			Str(XRequestID, xRequestId).
			Str(XForwardedFor, xForwardedFor).
			Str(TrueClientIP, trueClientIP).
			Str(CFConnectingIP, cfConnectingIP).
			Logger()
		if reqQueries != "" {
			zrLogger = zrLogger.With().Str("queries", reqQueries).Logger()
		}

		c.Next()

		l := &logger{
			GinCtx: c,
		}

		for _, key := range []string{ErrKey, InfoKey, DebugKey, CustomDataKey} {
			zrLogger = prepareLogData(l, zrLogger, key, defaultLogLvl)
		}

		var (
			finishTime        = time.Now()
			latency           = finishTime.Sub(startTime)
			statusCode        = c.Writer.Status()
			response, existed = c.Get(ContextResponse)
		)

		if existed {
			zrLogger.Log().
				Str("severity", getSeverityLevelLog(l, defaultLogLvl)).
				Dur("latency", latency).
				Int("status_code", statusCode).
				RawJSON("response", response.([]byte)).
				Str("message", fmt.Sprintf("[%s] %13v %d %s", reqMethod, latency, statusCode, reqEndpoint)).
				Msg(handlerName)
		} else {
			if statusCode > 399 && statusCode != 401 {
				zrLogger.Log().
					Str("user_agent", fmt.Sprintf("%s", c.Request.UserAgent())).
					Str("x_user_agent", fmt.Sprintf("%s", c.Request.Header.Get("X-User-Agent"))).
					Str("severity", getSeverityLevelLog(l, defaultLogLvl)).
					Dur("latency", latency).
					Int("status_code", statusCode).
					Str("message", fmt.Sprintf("[%s] %13v %d %s", reqMethod, latency, statusCode, reqEndpoint)).
					Msg(handlerName)
			} else {
				zrLogger.Log().
					Str("severity", getSeverityLevelLog(l, defaultLogLvl)).
					Str("user_agent", fmt.Sprintf("%s", c.Request.UserAgent())).
					Str("x_user_agent", fmt.Sprintf("%s", c.Request.Header.Get("X-User-Agent"))).
					Dur("latency", latency).
					Int("status_code", statusCode).
					Str("message", fmt.Sprintf("[%s] %13v %d %s", reqMethod, latency, statusCode, reqEndpoint)).
					Msg(handlerName)
			}
		}
	}
}

func getSeverityLevelLog(logger *logger, logLevelDefault zerolog.Level) string {
	var currentLogLevel = logLevelDefault
	if logData := logger.GetLogData(DebugKey); len(logData) != 0 && zerolog.DebugLevel >= logLevelDefault {
		currentLogLevel = zerolog.DebugLevel
	}

	if logData := logger.GetLogData(InfoKey); len(logData) != 0 && zerolog.InfoLevel >= logLevelDefault {
		currentLogLevel = zerolog.InfoLevel
	}

	if logData := logger.GetLogData(WarningKey); len(logData) != 0 && zerolog.WarnLevel >= logLevelDefault {
		currentLogLevel = zerolog.WarnLevel
	}

	if logData := logger.GetLogData(ErrKey); len(logData) != 0 && zerolog.ErrorLevel >= logLevelDefault {
		currentLogLevel = zerolog.ErrorLevel
	}

	switch currentLogLevel {
	case zerolog.DebugLevel:
		return DebugLevel
	case zerolog.InfoLevel:
		return InfoLevel
	case zerolog.WarnLevel:
		return WarningLevel
	case zerolog.ErrorLevel:
		return ErrorLevel
	default:
		return "DEFAULT"
	}
}

func prepareLogData(logger *logger, zrLogger zerolog.Logger, key string, currentLogLevel zerolog.Level) zerolog.Logger {
	logData := logger.GetLogData(key)
	switch key {
	case CustomDataKey:
		for _, key := range logData {
			zrLogger = zrLogger.With().Interface(key, logger.Get(key)).Logger()
		}
	case ErrKey:
		if zerolog.ErrorLevel >= currentLogLevel {
			for i, data := range logData {
				zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
			}
		}
	case InfoKey:
		if zerolog.InfoLevel >= currentLogLevel {
			for i, data := range logData {
				zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
			}
		}
	case DebugKey:
		if zerolog.DebugLevel >= currentLogLevel {
			for i, data := range logData {
				zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
			}
		}
	case WarningKey:
		if zerolog.WarnLevel >= currentLogLevel {
			for i, data := range logData {
				zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
			}
		}
	default:
		for i, data := range logData {
			zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
		}
	}
	return zrLogger
}
