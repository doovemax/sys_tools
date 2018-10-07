package log

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

type logger struct {
	// 日志传输管道
	Out chan *message
	// 管道缓冲容量
	MaxCache int
	// 日志等级
	LogLevel Level
	// 日志格式
	Format []Formater
	// 是否开启颜色显示,Default true
	EnableColor bool
	// 开启日志输出到文件,Default false
	EnableFileLog bool
	// 日志文件目录
	LogPath string
	// 日志文件名称
	LogFileName string

	Lock sync.Locker

	LogHandle
}

func (l *logger) SetMaxCache(n int) (err error) {
	l.Lock.Lock()
	defer l.Lock.Unlock()

	l.Out = make(chan *message, n)
	return

}

func (l *logger) Panic(m string) (err error) {
	l.Out <- &message{
		LogLevel: PanicLevel,
		Time:     time.Now(),
		Msg:      m,
	}
	return

}

func (l *logger) Fatal(m string) (err error) {
	l.Out <- &message{
		LogLevel: FatalLevel,
		Time:     time.Now(),
		Msg:      m,
	}
	return
}

func (l *logger) Error(m string) (err error) {
	l.Out <- &message{
		LogLevel: ErrorLevel,
		Time:     time.Now(),
		Msg:      m,
	}
	return
}
func (l *logger) Warn(m string) (err error) {
	l.Out <- &message{
		LogLevel: WarnLevel,
		Time:     time.Now(),
		Msg:      m,
	}
	return
}
func (l *logger) Info(m string) (err error) {
	l.Out <- &message{
		LogLevel: InfoLevel,
		Time:     time.Now(),
		Msg:      m,
	}
	return
}
func (l *logger) Debug(m string) (err error) {
	l.Out <- &message{
		Time:     time.Now(),
		LogLevel: DebugLevel,
		Msg:      m,
	}
	return
}

// 日志消息体
type message struct {
	LogLevel Level
	Time     time.Time
	Msg      string
}
