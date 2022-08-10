package log

import (
	"context"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

const (
	RequestIdKey = "x-amzn-RequestId"
)

var (
	severityMap = map[logrus.Level]byte{
		logrus.TraceLevel: 'D',
		logrus.DebugLevel: 'D',
		logrus.InfoLevel:  'I',
		logrus.WarnLevel:  'W',
		logrus.ErrorLevel: 'E',
		logrus.FatalLevel: 'F',
		logrus.PanicLevel: 'F',
	}
)

type formatter struct{}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Buffer.WriteByte('[')

	entry.Buffer.WriteByte(severityMap[entry.Level])

	entry.Buffer.WriteByte(' ')

	entry.Buffer.WriteString(entry.Time.Format("2006-01-02 15:04:05.000"))

	entry.Buffer.WriteByte(' ')

	var requestId string
	if entry.Context != nil {
		requestId, _ = entry.Context.Value(RequestIdKey).(string)
	}
	if requestId != "" {
		entry.Buffer.WriteString(requestId)
	} else {
		entry.Buffer.WriteByte('-')
	}

	entry.Buffer.WriteString("] ")

	if len(entry.Data) > 0 {
		if extraData, err := json.Marshal(entry.Data); err == nil {
			entry.Buffer.WriteString("[")
			entry.Buffer.Write(extraData)
			entry.Buffer.WriteString("] ")
		}
	}

	entry.Buffer.WriteString(entry.Message)

	entry.Buffer.WriteByte('\n')

	return entry.Buffer.Bytes(), nil
}

// fork from https://github.com/sirupsen/logrus/blob/v1.6.0/exported.go

var std *logrus.Logger

type Level = logrus.Level
type Fields = logrus.Fields

const (
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = logrus.ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = logrus.WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = logrus.InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = logrus.DebugLevel
)

func init() {
	std = &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(formatter),
		Hooks:        make(logrus.LevelHooks, 0),
		ReportCaller: false,
		Level:        logrus.InfoLevel,
	}
}

func StandardLogger() *logrus.Logger {
	return std
}

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
	std.SetLevel(level)
}

// GetLevel returns the standard logger level.
func GetLevel() Level {
	return std.GetLevel()
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level param
func IsLevelEnabled(level Level) bool {
	return std.IsLevelEnabled(level)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(ctx context.Context, fields Fields) *logrus.Entry {
	return std.WithContext(ctx).WithFields(fields)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(ctx context.Context, args ...interface{}) {
	std.WithContext(ctx).Debug(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(ctx context.Context, args ...interface{}) {
	std.WithContext(ctx).Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(ctx context.Context, args ...interface{}) {
	std.WithContext(ctx).Warn(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(ctx context.Context, args ...interface{}) {
	std.WithContext(ctx).Error(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	std.WithContext(ctx).Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(ctx context.Context, format string, args ...interface{}) {
	std.WithContext(ctx).Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(ctx context.Context, format string, args ...interface{}) {
	std.WithContext(ctx).Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	std.WithContext(ctx).Errorf(format, args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(ctx context.Context, args ...interface{}) {
	std.WithContext(ctx).Debugln(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(ctx context.Context, args ...interface{}) {
	std.WithContext(ctx).Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(ctx context.Context, args ...interface{}) {
	std.WithContext(ctx).Warnln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(ctx context.Context, args ...interface{}) {
	std.WithContext(ctx).Errorln(args...)
}

func RequestIdFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	val := ctx.Value(RequestIdKey)
	if sp, ok := val.(string); ok {
		return sp
	}
	return ""
}

func ContextWithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, RequestIdKey, requestId)
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := uuid.NewV4()
		id := uid.String()
		c.Set(RequestIdKey, id)
		c.Header(RequestIdKey, id)
		Info(c, "request start")
		c.Next()
		Info(c, "request end")
	}
}
