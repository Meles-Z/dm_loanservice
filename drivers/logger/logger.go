package logger

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

var logger log.Logger

func InitLogger() {
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.Caller(4))
}

func GetLogger() log.Logger {
	return logger
}

func LogError(keyvals ...interface{}) {
	_ = level.Error(logger).Log(keyvals...)
}

func LogInfo(keyvals ...interface{}) {
	_ = level.Info(logger).Log(keyvals...)
}

func LogDebug(keyvals ...interface{}) {
	_ = level.Debug(logger).Log(keyvals...)
}
