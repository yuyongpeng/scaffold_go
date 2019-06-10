package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"scaffold_go/conf"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// 设置日志的输出
	switch conf.LOG_OUTPUT {
	case "file":
		file, err := os.OpenFile(conf.LOG_OUTPUT_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			Log.Out = file
		} else {
			Log.Out = os.Stderr
			Log.Error("Failed to log to file, using default stderr")
		}
	case "console":
		Log.Out = os.Stderr
	default:
		Log.Out = os.Stderr
	}

	// 设置日志的格式
	switch conf.LOG_FORMATER {
	case "json":
		Log.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		Log.SetFormatter(&logrus.TextFormatter{})
	default:
		Log.SetFormatter(&logrus.TextFormatter{})
	}

	// 设置日志的输出级别
	switch conf.LOG_LEVEL {
	case "trace":
		Log.SetLevel(logrus.TraceLevel)
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		Log.SetLevel(logrus.FatalLevel)
	case "panic":
		Log.SetLevel(logrus.PanicLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}
}
