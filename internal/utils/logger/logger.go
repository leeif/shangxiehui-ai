package logger

import (
	"os"
	"shangxiehui-ai/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type KiwiLogger struct {
	DefaultLogger  *zap.Logger
	LLMChainLogger *zap.Logger
}

func (log *KiwiLogger) With(fields ...zapcore.Field) *KiwiLogger {
	newDefaultLogger := log.DefaultLogger.With(fields...)

	return &KiwiLogger{
		DefaultLogger:  newDefaultLogger,
		LLMChainLogger: log.LLMChainLogger,
	}
}

func (log *KiwiLogger) Debug(msg string, fields ...zapcore.Field) {
	log.DefaultLogger.Debug(msg, fields...)
}

func (log *KiwiLogger) Info(msg string, fields ...zapcore.Field) {
	log.DefaultLogger.Info(msg, fields...)
}

func (log *KiwiLogger) Warn(msg string, fields ...zapcore.Field) {
	log.DefaultLogger.Warn(msg, fields...)
}

func (log *KiwiLogger) Error(msg string, fields ...zapcore.Field) {
	log.DefaultLogger.Error(msg, fields...)
}

func (log *KiwiLogger) DPanic(msg string, fields ...zapcore.Field) {
	log.DefaultLogger.DPanic(msg, fields...)
}

func (log *KiwiLogger) Panic(msg string, fields ...zapcore.Field) {
	log.DefaultLogger.Panic(msg, fields...)
}

func (log *KiwiLogger) Fatal(msg string, fields ...zapcore.Field) {
	log.DefaultLogger.Fatal(msg, fields...)
}

func (log *KiwiLogger) LLMChain(chain, prompt, response string) {
	log.LLMChainLogger.Info(chain, zap.String("prompt", prompt), zap.String("response", response))
}

func (log *KiwiLogger) Sync() error {
	log.DefaultLogger.Sync()
	log.LLMChainLogger.Sync()
	return nil
}

func newLLMChainLogger(config *config.Config) (*zap.Logger, error) {

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(pe)

	var syncer zapcore.WriteSyncer
	if config.Log.LLMChainFile == "" {
		syncer = zapcore.AddSync(os.Stdout)
	} else {
		syncer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.Log.LLMChainFile,
			MaxSize:    500,
			MaxBackups: 100,
			MaxAge:     7,
			Compress:   true,
		})
	}

	level := zap.InfoLevel

	core := zapcore.NewCore(encoder, syncer, zap.NewAtomicLevelAt(level))

	logger := zap.New(core, zap.WithCaller(true))

	return logger, nil
}

func newDefaultLogger(config *config.Config) (*zap.Logger, error) {

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder

	encoder := zapcore.NewJSONEncoder(pe)

	switch config.Log.Format {
	case "json":
		encoder = zapcore.NewJSONEncoder(pe)
	case "console":
		encoder = zapcore.NewConsoleEncoder(pe)
	}

	var syncer zapcore.WriteSyncer
	if config.Log.File != "" {
		syncer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.Log.File,
			MaxSize:    500,
			MaxBackups: 100,
			MaxAge:     7,
			Compress:   true,
		})
	} else {
		syncer = zapcore.AddSync(os.Stdout)
	}

	level := zap.InfoLevel

	switch config.Log.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	case "dev":
		level = zap.DPanicLevel
	}

	core := zapcore.NewCore(encoder, syncer, zap.NewAtomicLevelAt(level))

	logger := zap.New(core, zap.WithCaller(true))

	zap.ReplaceGlobals(logger)

	return logger, nil
}

func NewKiwiLogger(cfg *config.Config) (*KiwiLogger, error) {
	defaultLogger, err := newDefaultLogger(cfg)
	if err != nil {
		return nil, err
	}

	llmChainLogger, err := newLLMChainLogger(cfg)
	if err != nil {
		return nil, err
	}

	return &KiwiLogger{
		DefaultLogger:  defaultLogger,
		LLMChainLogger: llmChainLogger,
	}, nil
}
