package clog

import (
	"sync"
	"time"
)

type Level uint

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

type Logger struct {
	// 日志传输管道
	Out chan *message
	// 管道缓冲容量
	MaxCache int
	// 日志等级
	LogLevel Level
	// 日志格式
	Format []Formater

	// 开启日志输出到文件,Default false
	EnableFileLog bool
	// 日志文件目录
	LogPath string
	// 日志文件名称
	LogFileName string

	Lock     sync.Locker
	Timezone string
}

func (l *Logger) SetMaxCache(n int) (err error) {
	l.Lock.Lock()
	defer l.Lock.Unlock()

	l.Out = make(chan *message, n)
	return

}

func (l *Logger) Run() (err error) {
	l.Lock.Lock()
	l.Out = make(chan *message, l.MaxCache)
	l.Lock.Unlock()
	go logout(l)
	return
}

func (l *Logger) Panic(v ...interface{}) (err error) {
	panic(v)
	return

}

func (l *Logger) Fatal(v ...interface{}) (err error) {
	l.Out <- &message{
		LogLevel: FatalLevel,
		Time:     time.Now(),
		Msg:      v,
	}
	return
}

func (l *Logger) Error(v ...interface{}) (err error) {
	l.Out <- &message{
		LogLevel: ErrorLevel,
		Time:     time.Now(),
		Msg:      v,
	}
	return
}
func (l *Logger) Warn(v ...interface{}) (err error) {
	l.Out <- &message{
		LogLevel: WarnLevel,
		Time:     time.Now(),
		Msg:      v,
	}
	return
}
func (l *Logger) Info(v ...interface{}) (err error) {
	l.Out <- &message{
		LogLevel: InfoLevel,
		Time:     time.Now(),
		Msg:      v,
	}
	return
}
func (l *Logger) Debug(v ...interface{}) (err error) {
	l.Out <- &message{
		Time:     time.Now(),
		LogLevel: DebugLevel,
		Msg:      v,
	}
	return
}

// 日志消息体
type message struct {
	LogLevel Level
	Time     time.Time
	Msg      interface{}
}
