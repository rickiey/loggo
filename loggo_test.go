package loggo

import (
	"errors"
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
