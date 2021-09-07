package loggo

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"strings"
)

type zapLog struct {
	*zap.SugaredLogger
}

func NewZapLog(level string, output io.Writer) Log {
	alevel := zap.NewAtomicLevel()
	w := zapcore.AddSync(output)

	alevel.SetLevel(ParseZapLevel(level))
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.CallerKey = "file"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		alevel,
	)
	return &zapLog{zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2)).Sugar()}
}

func (z *zapLog) Print(v ...interface{}) {
	z.SugaredLogger.Info(v...)
}

func (z *zapLog) Println(v ...interface{}) {
	z.SugaredLogger.Info(v...)
}

func (z *zapLog) Debug(v ...interface{}) {
	z.SugaredLogger.Debug(v...)
}

func (z *zapLog) Debugf(msg string, v ...interface{}) {
	z.SugaredLogger.Debugf(msg, v...)
}

func (z *zapLog) Info(v ...interface{}) {
	z.SugaredLogger.Info(v...)
}

func (z *zapLog) Infof(msg string, v ...interface{}) {
	z.SugaredLogger.Infof(msg, v...)
}

func (z *zapLog) Warn(v ...interface{}) {
	z.SugaredLogger.Warn(v...)
}

func (z *zapLog) Warnf(msg string, v ...interface{}) {
	z.SugaredLogger.Warnf(msg, v...)
}

func (z *zapLog) Error(v ...interface{}) {
	z.SugaredLogger.Error(v...)
}

func (z *zapLog) Errorf(msg string, v ...interface{}) {
	z.SugaredLogger.Errorf(msg, v...)
}

func (z *zapLog) Panic(v ...interface{}) {
	z.SugaredLogger.Panic(v...)
}

func (z *zapLog) Panicf(msg string, v ...interface{}) {
	z.SugaredLogger.Panicf(msg, v...)
}

func (z *zapLog) Fatal(v ...interface{}) {
	z.SugaredLogger.Fatal(v...)
}

func (z *zapLog) Fatalf(msg string, v ...interface{}) {
	z.SugaredLogger.Fatalf(msg, v...)
}

func ParseZapLevel(text string) zapcore.Level {

	l := zapcore.InfoLevel
	switch strings.TrimSpace(text) {
	case "debug", "DEBUG":
		l = zapcore.DebugLevel
	case "info", "INFO", "": // make the zero value useful
		l = zapcore.InfoLevel
	case "warn", "warning", "WARN":
		l = zapcore.WarnLevel
	case "error", "ERROR":
		l = zapcore.ErrorLevel
	case "panic", "PANIC":
		l = zapcore.PanicLevel
	case "fatal", "FATAL":
		l = zapcore.FatalLevel
	}
	return l
}
