package log

import (
	"sync"
)

var (
	Clog = New()
)

func init() {
	Clog.Out = make(chan message, Clog.MaxCache)
}

func New() *logger {
	return &logger{
		Out:           nil,
		MaxCache:      1000,
		LogLevel:      InfoLevel,
		Format:        nil,
		EnableColor:   true,
		EnableFileLog: false,
		LogPath:       "/tmp/",
		LogFileName:   "clog.log",
		Lock:          &sync.Mutex{},
	}
}
