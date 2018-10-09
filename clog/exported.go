package clog

import (
	"log"
	"sync"
)

var (
	Clog = New()
)

func init() {
	Clog.Out = make(chan *message, Clog.MaxCache)
	err := Clog.Run()
	if err != nil {
		log.Panic(err)
	}
}

func New() *Logger {
	return &Logger{
		Out:           nil,
		MaxCache:      1000,
		LogLevel:      InfoLevel,
		Format:        nil,
		EnableFileLog: false,
		LogPath:       "/tmp/",
		LogFileName:   "clog.log",
		Lock:          &sync.Mutex{},
		Timezone:      "Local",
	}
}
