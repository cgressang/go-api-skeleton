package zerolog

import (
	"fmt"
	"os"

	"github.com/cgressang/go-api-skeleton/usecase/log"
	"github.com/rs/zerolog"
)

type Logger struct {
	Log zerolog.Logger
}

func New(debug bool) log.StdLogger {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return &Logger{
		Log: zerolog.New(os.Stdout),
	}
}

func (l Logger) Debugf(format string, args ...interface{}) {
	l.Log.Debug().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.Log.Error().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	l.Log.Fatal().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Infof(format string, args ...interface{}) {
	l.Log.Info().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Panicf(format string, args ...interface{}) {
	l.Log.Panic().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l Logger) Warningf(format string, args ...interface{}) {
	l.Log.Warn().Msg(fmt.Sprintf(format, args...))
}

func (l Logger) Debug(args ...interface{}) {
	l.Log.Debug().Msg(fmt.Sprint(args...))
}

func (l Logger) Error(args ...interface{}) {
	l.Log.Error().Msg(fmt.Sprint(args...))
}

func (l Logger) Fatal(args ...interface{}) {
	l.Log.Fatal().Msg(fmt.Sprint(args...))
}

func (l Logger) Info(args ...interface{}) {
	l.Log.Info().Msg(fmt.Sprint(args...))
}

func (l Logger) Panic(args ...interface{}) {
	l.Log.Panic().Msg(fmt.Sprint(args...))
}

func (l Logger) Print(args ...interface{}) {
	l.Info(args...)
}

func (l Logger) Warning(args ...interface{}) {
	l.Log.Warn().Msg(fmt.Sprint(args...))
}

func (l Logger) Debugln(args ...interface{}) {
	l.Debug(args...)
}

func (l Logger) Errorln(args ...interface{}) {
	l.Error(args...)
}

func (l Logger) Fatalln(args ...interface{}) {
	l.Fatal(args...)
}

func (l Logger) Infoln(args ...interface{}) {
	l.Info(args...)
}

func (l Logger) Panicln(args ...interface{}) {
	l.Panic(args...)
}

func (l Logger) Println(args ...interface{}) {
	l.Info(args...)
}

func (l Logger) Warningln(args ...interface{}) {
	l.Warning(args...)
}
