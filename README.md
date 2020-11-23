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
}
```

+ 提供了 zap 和 logrus 两个实现，默认为 zap 实现，用户也可以自己实现并通过 SetLog 替换默认实现