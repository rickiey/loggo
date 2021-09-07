package loggo

import (
	"github.com/sirupsen/logrus"
	"io"
)

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
