package loggo

import (
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
)

func init() {
	Logger = NewZapLog("info", os.Stdout)
}

var Logger Log

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
	Fatal(v ...interface{})
	Fatalf(msg string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
}

func Debug(v ...interface{})              { Logger.Debug(v...) }
func Debugf(msg string, v ...interface{}) { Logger.Debugf(msg, v...) }
func Info(v ...interface{})               { Logger.Info(v...) }
func Infof(msg string, v ...interface{})  { Logger.Infof(msg, v...) }
func Warn(v ...interface{})               { Logger.Warn(v...) }
func Warnf(msg string, v ...interface{})  { Logger.Warnf(msg, v...) }
func Error(v ...interface{})              { Logger.Error(v...) }
func Errorf(msg string, v ...interface{}) { Logger.Errorf(msg, v...) }
func Panic(v ...interface{})              { Logger.Panic(v...) }
func Panicf(msg string, v ...interface{}) { Logger.Panicf(msg, v...) }
func Fatal(v ...interface{})              { Logger.Fatal(v...) }
func Fatalf(msg string, v ...interface{}) { Logger.Fatalf(msg, v...) }

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
	Logger = l
}

type logrusLog struct{}

func (l *logrusLog) Print(v ...interface{}) {
	logrus.Print(v...)
}

func (l *logrusLog) Println(v ...interface{}) {
	logrus.Println(v...)
}

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

func (l *logrusLog) Fatal(v ...interface{}) {
	logrus.Fatal(v...)
}

func (l *logrusLog) Fatalf(msg string, v ...interface{}) {
	logrus.Fatalf(msg, v...)
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

func NewZerolog(level string, output io.Writer) Log {

	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	log := zerolog.New(output).With().Caller().Timestamp().Logger()
	log.Level(lvl)

	return &ZeloLog{&log}
}

type ZeloLog struct {
	l *zerolog.Logger
}

func interfaces(v ...interface{}) []interface{} {
	return append([]interface{}{}, v...)
}

func (z *ZeloLog) Debug(v ...interface{}) {
	z.l.Debug().Interface("", interfaces(v)).Send()
}

func (z *ZeloLog) Debugf(msg string, v ...interface{}) {
	z.l.Debug().Msgf(msg, v...)
}

func (z *ZeloLog) Info(v ...interface{}) {
	z.l.Info().Interface("", interfaces(v)).Send()
}

func (z *ZeloLog) Infof(msg string, v ...interface{}) {
	z.l.Info().Msgf(msg, v...)
}

func (z *ZeloLog) Warn(v ...interface{}) {
	z.l.Warn().Interface("", interfaces(v)).Send()
}

func (z ZeloLog) Warnf(msg string, v ...interface{}) {
	z.l.Warn().Msgf(msg, v...)
}

func (z *ZeloLog) Error(v ...interface{}) {
	z.l.Debug().Interface("", interfaces(v)).Send()
}

func (z *ZeloLog) Errorf(msg string, v ...interface{}) {
	z.l.Error().Msgf(msg, v...)
}

func (z *ZeloLog) Panic(v ...interface{}) {
	z.l.Panic().Interface("", interfaces(v)).Send()
}

func (z *ZeloLog) Panicf(msg string, v ...interface{}) {
	z.l.Panic().Msgf(msg, v...)
}

func (z *ZeloLog) Fatal(v ...interface{}) {
	z.l.Fatal().Interface("", interfaces(v)).Send()
}

func (z *ZeloLog) Fatalf(msg string, v ...interface{}) {
	z.l.Fatal().Msgf(msg, v...)
}

func (z *ZeloLog) Print(v ...interface{}) {
	z.l.Print(v...)
}

func (z *ZeloLog) Println(v ...interface{}) {
	z.l.Print(v...)
}
