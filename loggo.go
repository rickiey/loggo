package loggo

import (
	"os"
)

func init() {
	Logger = NewZapLog("info", os.Stdout)
}

var Logger Log

func SetLog(l Log) {
	Logger = l
}

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
