# loggo

### go 一个日志接口

+ 提供了一个日志接口

```go

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
```

+ 提供了 zap 和 logrus 两个实现，默认为 zap 实现，用户也可以自己实现并通过 SetLog 替换默认实现

> 使用示例：

```go
func ConfigLog() {
	output := &lumberjack.Logger{
		Filename: conf.CFG.LogFile,

		// default 100
		MaxSize: 300, // megabytes

		MaxBackups: 60,
		MaxAge:     60,   //days
		Compress:   true, // disabled by default
	}
	loggo.NewZapLog(conf.CFG.LogLevel, output)
	loggo.Info("loggo configuration done.")
}

```

> 自定义 log, 只需实现 Log 接口

```go
package main

import (
    "fmt"
    "github.com/rickiey/loggo"
)

func main() {
    loggo.SetLog(new(flog))
    loggo.Info("loggo ...")
}

type flog struct{}
func (l *flog) Print(v ...interface{}) { fmt.Println(v...) }
func (l *flog) Println(v ...interface{}) { fmt.Println(v...) }

func (l *flog) Debug(v ...interface{}) {fmt.Println(v...)}

func (l *flog) Debugf(msg string, v ...interface{}) {fmt.Printf(msg, v...)}

func (l *flog) Info(v ...interface{}) {fmt.Println(v...)}

func (l *flog) Infof(msg string, v ...interface{}) {fmt.Printf(msg, v...)}

func (l *flog) Warn(v ...interface{}) {fmt.Println(v...)}

func (l *flog) Warnf(msg string, v ...interface{}) {fmt.Printf(msg, v...)}

func (l *flog) Error(v ...interface{}) {fmt.Println(v...)}

func (l *flog) Errorf(msg string, v ...interface{}) {fmt.Printf(msg, v...)}

func (l *flog) Panic(v ...interface{}) {fmt.Println(v...)}

func (l *flog) Panicf(msg string, v ...interface{}) {fmt.Printf(msg, v...)}
func (l *flog) Fatal(v ...interface{}) { fmt.Println(v...) }
func (l *flog) Fatalf(msg string, v ...interface{}) { fmt.Printf(msg, v...) }
```