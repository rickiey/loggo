package loggo

import (
	"github.com/rs/zerolog"
	"io"
)

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
