package ezLogzero

import (
	"github.com/rs/zerolog"
)

type LogOption func(o *logOptions)

type logOptions struct {
	logFilePath string
	logFileName string

	logUnixTime   bool
	logLevel      zerolog.Level
	logWithCaller bool
	customTimeKey string
	customMsgKey  string

	globalPrefix map[string]interface{}

	prefixTimestampEnable bool
	prefixTimestampKey    string

	//rotate config
	logMaxSize        int  //日志文件最大尺寸(MB)
	logMaxAge         int  //日志最大保留天数,0为保留所有
	logMaxBackups     int  //日志文件保留份数,0为保留所有
	logBackupCompress bool //日志备份文件是否压缩存储
}

func WithOutputFilePath(path string, filename string) LogOption {
	return func(o *logOptions) {
		o.logFilePath = path
		o.logFileName = filename
	}
}

func WithOutputFileRotate(maxSize int, maxAge int, maxBackup int, compressEnable bool) LogOption {
	return func(o *logOptions) {
		o.logMaxSize = maxSize
		o.logMaxAge = maxAge
		o.logMaxBackups = maxBackup
		o.logBackupCompress = compressEnable
	}
}

func WithLogUnixTimestamp() LogOption {
	return func(o *logOptions) {
		o.logUnixTime = true
	}
}

func WithLevelTrace() LogOption {
	return func(o *logOptions) {
		o.logLevel = zerolog.TraceLevel
	}
}

func WithLevelDebug() LogOption {
	return func(o *logOptions) {
		o.logLevel = zerolog.DebugLevel
	}
}

func WithLevelInfo() LogOption {
	return func(o *logOptions) {
		o.logLevel = zerolog.InfoLevel
	}
}

func WithLevelWarn() LogOption {
	return func(o *logOptions) {
		o.logLevel = zerolog.WarnLevel
	}
}

func WithLevelError() LogOption {
	return func(o *logOptions) {
		o.logLevel = zerolog.ErrorLevel
	}
}

func WithLevelFatal() LogOption {
	return func(o *logOptions) {
		o.logLevel = zerolog.FatalLevel
	}
}

func WithLevelPanic() LogOption {
	return func(o *logOptions) {
		o.logLevel = zerolog.PanicLevel
	}
}

func WithGlobalPrefix(prefix map[string]interface{}) LogOption {
	return func(o *logOptions) {
		o.globalPrefix = prefix
	}
}

func WithPrefixTimestamp(k string) LogOption {
	return func(o *logOptions) {
		o.prefixTimestampEnable = true
		o.prefixTimestampKey = k
	}
}

func WithCaller() LogOption {
	return func(o *logOptions) {
		o.logWithCaller = true
	}
}

func WithCustomTimeKey(k string) LogOption {
	return func(o *logOptions) {
		o.customTimeKey = k
	}
}

func WithCustomMsgKey(k string) LogOption {
	return func(o *logOptions) {
		o.customMsgKey = k
	}
}
