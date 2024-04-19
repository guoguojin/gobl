package logger

import (
	"errors"
	"io"
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"gitlab.com/gobl/gobl/pkg/config"
)

//nolint:gochecknoglobals
var (
	once   sync.Once
	syncer zapcore.WriteSyncer
	core   zapcore.Core
	log    *zap.Logger
)

const (
	defaultLogFileName    = "./log/application.log"
	defaultMaxSizeMB      = 100
	defaultMaxBackupFiles = 5
	defaultMaxAgeDays     = 30
	defaultUseCompression = false
)

// ErrCoreNotInitialised is returned if the zap core has not been initialised
var ErrCoreNotInitialised = errors.New("zap core has not been initialised")

// ZapConfig returns a default configuration for the Zap core
func ZapConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return encoderConfig
}

// ZapEncoder() returns a console encoder using the default ZapConfig
func ZapEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(ZapConfig())
}

// ZapWriter returns a synchronous Zap writer writing to a log file and Stdout at
// the same time
func ZapWriter(writer io.Writer) zapcore.WriteSyncer {
	if syncer == nil {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer), zapcore.AddSync(os.Stdout))
	}

	return syncer
}

// LumberjackLogger returns a Lumberjack rotating log writer to be used with the ZapWriter
func LumberjackLogger(fileName string, maxSize, maxBackups, maxAge int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}

// ZapCore returns the zap core that is used by the logger
func ZapCore() (zapcore.Core, error) {
	if core == nil {
		return nil, ErrCoreNotInitialised
	}

	return core, nil
}

// Get returns a new or current configured zap logger
func Get(level zapcore.Level, writer io.Writer) *zap.Logger {
	once.Do(func() {
		core = zapcore.NewCore(ZapEncoder(), ZapWriter(writer), level)
		log = zap.New(core, zap.AddCaller())
	})

	defer func() {
		_ = log.Sync()
	}()

	return log
}

// New creates a new zap logger using the default encoder and provided writer
func New(level zapcore.Level, writer io.Writer) *zap.Logger {
	core := zapcore.NewCore(ZapEncoder(), ZapWriter(writer), level)
	return zap.New(core, zap.AddCaller())
}

// ConfiguredLumberjackLogger returns a lumberjack logger using configuration provided
// in the application configuration file or using predefined defaults if they are not
// available:
// Logfile: ./log/application.log
// Max File Size: 100MB
// Max Backups: 5
// Max Age of log files: 30 days
// Use compression: false
func ConfiguredLumberjackLogger() *lumberjack.Logger {
	return LumberjackLogger(
		config.Get(config.LogFilePathKey).String(defaultLogFileName),
		config.Get(config.LogFileMaxSize).Int(defaultMaxSizeMB),
		config.Get(config.LogFileMaxBackups).Int(defaultMaxBackupFiles),
		config.Get(config.LogFileMaxAge).Int(defaultMaxAgeDays),
		config.Get(config.LogFileCompress).Bool(defaultUseCompression),
	)
}

// ConsoleLogger returns a zap logger that writes to Stdout by default at INFO level
func ConsoleLogger() *zap.Logger {
	c := zapcore.NewCore(ZapEncoder(), zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
	return zap.New(c, zap.AddCaller())
}

// ApplicationLogLevel returns the log level defined in the
// application configuration file
func ApplicationLogLevel() zapcore.Level {
	var level zapcore.Level

	switch strings.ToUpper(config.Get(config.LogLevelKey).String("")) {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	case "FATAL":
		level = zapcore.FatalLevel
	case "PANIC":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	return level
}

// Logger logger will create a new logger with the configured application log level
// and lumberjack logger if one doesn't exist, or return the created logger if it has
func Logger() *zap.Logger {
	return Get(ApplicationLogLevel(), ConfiguredLumberjackLogger())
}
