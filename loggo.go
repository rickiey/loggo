package loggo

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
)

func init() {
	logger = NewZapLog("info", os.Stdout)
}

var logger Log

type Log interface {
	Debug(v ...interface{})
	Debugf(msg string, v ...interface{})
	Info(v ...interface{})
	Infof(msg string, v ...interface{})
	Warn(v ...interface{})
	Warnf(msg string, v ...interface{})
	Error(v ...interface{})
	Errorf(msg string, v ...interface{})
	Panic(v ...interface{})
	Panicf(msg string, v ...interface{})
}

func Debug(v ...interface{})              { logger.Debug(v...) }
func Debugf(msg string, v ...interface{}) { logger.Debugf(msg, v...) }
func Info(v ...interface{})               { logger.Info(v...) }
func Infof(msg string, v ...interface{})  { logger.Infof(msg, v...) }
func Warn(v ...interface{})               { logger.Warn(v...) }
func Warnf(msg string, v ...interface{})  { logger.Warnf(msg, v...) }
func Error(v ...interface{})              { logger.Error(v...) }
func Errorf(msg string, v ...interface{}) { logger.Errorf(msg, v...) }
func Panic(v ...interface{})              { logger.Panic(v...) }
func Panicf(msg string, v ...interface{}) { logger.Panicf(msg, v...) }

func NewLogrusLog(level string, output io.Writer) Log {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	defaultLevel := logrus.InfoLevel
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(defaultLevel)
		logrus.Warn("logrus level error :", err)
	}

	logrus.SetOutput(output)
	logrus.SetLevel(lv)
	return new(logrusLog)
}

func SetLog(l Log) {
	logger = l
}

type logrusLog struct{}

func (l *logrusLog) Debug(v ...interface{}) {
	logrus.Debug(v...)
}

func (l *logrusLog) Debugf(msg string, v ...interface{}) {
	logrus.Debugf(msg, v...)
}

func (l *logrusLog) Info(v ...interface{}) {
	logrus.Info(v...)
}

func (l *logrusLog) Infof(msg string, v ...interface{}) {
	logrus.Infof(msg, v...)
}

func (l *logrusLog) Warn(v ...interface{}) {
	logrus.Warn(v...)
}

func (l *logrusLog) Warnf(msg string, v ...interface{}) {
	logrus.Warnf(msg, v...)
}

func (l *logrusLog) Error(v ...interface{}) {
	logrus.Error(v...)
}

func (l *logrusLog) Errorf(msg string, v ...interface{}) {
	logrus.Errorf(msg, v...)
}

func (l *logrusLog) Panic(v ...interface{}) {
	logrus.Panic(v...)
}

func (l *logrusLog) Panicf(msg string, v ...interface{}) {
	logrus.Panicf(msg, v...)
}

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
	return &zapLog{zap.New(core, zap.AddCaller()).Sugar()}
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
