package loggo

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func Test_zaplog(t *testing.T) {
	l := NewZapLog("debug", os.Stdout)
	SetLog(l)
	Debug("1", 1)
	Debugf("%v %v", "1", 1)
	Info("1", 1)
	Infof("%v %v", "1", 1)
	Warn("1", 1)
	Warnf("%v %v", "1", 1)
	Error("1", 1)
	Errorf("%v %v", "1", 1)
	//Panic("1", 1)
	//Panicf("%v %v", "1", 1)
}

func Test_logrus(t *testing.T) {
	l := NewLogrusLog("debug", os.Stdout)
	SetLog(l)
	Debug(1, 2, 3, "666", errors.New("errormsg"))
	Debugf("%v %v, %v, %v, %v", 1, 2, 3, "666", errors.New("errormsg"))
	Info(1, 2, 3, "666", errors.New("errormsg"))
	Infof("%v %v, %v, %v, %v", 1, 2, 3, "666", errors.New("errormsg"))
	Warn(1, 2, 3, "666", errors.New("errormsg"))
	Warnf("%v %v, %v, %v, %v", 1, 2, 3, "666", errors.New("errormsg"))
	Error(1, 2, 3, "666", errors.New("errormsg"))
	Errorf("%v %v, %v, %v, %v", 1, 2, 3, "666", errors.New("errormsg"))
	//Panic(1, 2, 3, "666", errors.New("errormsg"))
	//Panicf("%v %v, %v, %v, %v", 1, 2, 3, "666", errors.New("errormsg"))
}

func Test_fmt(t *testing.T) {
	SetLog(new(flog))
	Debug("1", 1)
	Debugf("%v %v", "1", 1)
	Info("1", 1)
	Infof("%v %v", "1", 1)
	Warn("1", 1)
	Warnf("%v %v", "1", 1)
	Error("1", 1)
	Errorf("%v %v", "1", 1)
}

type flog struct{}

func (l *flog) Debug(v ...interface{}) { fmt.Println(v...) }

func (l *flog) Debugf(msg string, v ...interface{}) { fmt.Printf(msg, v...) }

func (l *flog) Info(v ...interface{}) { fmt.Println(v...) }

func (l *flog) Infof(msg string, v ...interface{}) { fmt.Printf(msg, v...) }

func (l *flog) Warn(v ...interface{}) { fmt.Println(v...) }

func (l *flog) Warnf(msg string, v ...interface{}) { fmt.Printf(msg, v...) }

func (l *flog) Error(v ...interface{}) { fmt.Println(v...) }

func (l *flog) Errorf(msg string, v ...interface{}) { fmt.Printf(msg, v...) }

func (l *flog) Panic(v ...interface{}) { fmt.Println(v...) }

func (l *flog) Panicf(msg string, v ...interface{}) { fmt.Printf(msg, v...) }
