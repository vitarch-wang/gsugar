package ezLogzero

import (
	"fmt"
	"github.com/rs/zerolog"
	"golang.org/x/net/context"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
	"time"
)

type Log struct {
	logger   *zerolog.Logger
	opts     logOptions
	ctx      context.Context
	initTime time.Time
}

func NewLogger(opts ...LogOption) (*Log, error) {

	//init options
	options := logOptions{
		logLevel:          zerolog.InfoLevel,
		logFilePath:       "./logs/",
		logFileName:       "access",
		logMaxSize:        10,
		logMaxAge:         1,
		logMaxBackups:     1,
		logBackupCompress: true,
	}

	//reset options
	for _, o := range opts {
		o(&options)
	}

	zerolog.SetGlobalLevel(options.logLevel)

	if options.logUnixTime {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	if len(options.customTimeKey) > 0 {
		zerolog.TimestampFieldName = options.customTimeKey
	}

	if len(options.customMsgKey) > 0 {
		zerolog.MessageFieldName = options.customMsgKey
	}

	rotateWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", options.logFilePath, options.logFileName), //filePath
		MaxSize:    options.logMaxSize,                                                 // megabytes
		MaxBackups: options.logMaxBackups,
		MaxAge:     options.logMaxAge,         //days
		Compress:   options.logBackupCompress, // disabled by default
	}

	logger := zerolog.New(rotateWriter)
	if len(options.globalPrefix) > 0 {
		for k, v := range options.globalPrefix {
			logger = logger.With().Interface(k, v).Logger()
		}
	}
	if options.logWithCaller {
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			var fileName string
			pathSplit := strings.Split(file, "/")
			fileDir := ""
			if len(pathSplit) > 1 {
				fileDir = pathSplit[len(pathSplit)-2] + "/"
			}
			fileName = pathSplit[len(pathSplit)-1]
			return fmt.Sprintf("%s%s:%d", fileDir, fileName, line)
		}
		logger = logger.With().Caller().Logger()
	}

	return &Log{opts: options, logger: &logger}, nil
}

func (l *Log) Writer() *zerolog.Logger {
	return l.logger
}

func (l *Log) logSub(prefix map[string]interface{}) *zerolog.Logger {
	subLogger := l.logger.With().Logger()
	if len(prefix) > 0 {
		for k, v := range prefix {
			subLogger = subLogger.With().Interface(k, v).Logger()
		}
	}
	if l.opts.prefixTimestampEnable {
		subLogger = subLogger.With().Int64(l.opts.prefixTimestampKey, time.Now().Unix()).Logger()
	}

	return &subLogger
}

func (l *Log) Info() *zerolog.Event {
	return l.logSub(nil).Info()
}

func (l *Log) Error() *zerolog.Event {
	return l.logSub(nil).Error()
}

func (l *Log) Warn() *zerolog.Event {
	return l.logSub(nil).Warn()
}

func (l *Log) Debug() *zerolog.Event {
	return l.logSub(nil).Debug()
}

func (l *Log) Trace() *zerolog.Event {
	return l.logSub(nil).Trace()
}

func (l *Log) Fatal() *zerolog.Event {
	return l.logSub(nil).Fatal()
}

func (l *Log) Panic() *zerolog.Event {
	return l.logSub(nil).Panic()
}

func (l *Log) InfoP(prefix map[string]interface{}) *zerolog.Event {
	return l.logSub(prefix).Info()
}

func (l *Log) ErrorP(prefix map[string]interface{}) *zerolog.Event {
	return l.logSub(prefix).Error()
}

func (l *Log) WarnP(prefix map[string]interface{}) *zerolog.Event {
	return l.logSub(prefix).Warn()
}

func (l *Log) DebugP(prefix map[string]interface{}) *zerolog.Event {
	return l.logSub(prefix).Debug()
}

func (l *Log) TraceP(prefix map[string]interface{}) *zerolog.Event {
	return l.logSub(prefix).Trace()
}

func (l *Log) FatalP(prefix map[string]interface{}) *zerolog.Event {
	return l.logSub(prefix).Fatal()
}

func (l *Log) PanicP(prefix map[string]interface{}) *zerolog.Event {
	return l.logSub(prefix).Panic()
}
