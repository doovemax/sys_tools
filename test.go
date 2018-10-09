package main

import (
	"sync"
	"time"

	"github.com/doovemax/sys_tools/clog"
)

func main() {
	log := &clog.Logger{
		Out:           nil,
		MaxCache:      1000,
		LogLevel:      clog.DebugLevel,
		Format:        nil,
		EnableFileLog: true,
		LogPath:       "/tmp/",
		LogFileName:   "clog.log",
		Lock:          &sync.Mutex{},
		Timezone:      "Local",
	}
	log.Run()

	// log.Fatal("this is a test")
	log.Error("this is b test")
	log.Warn("this is c test")
	log.Info("this is d test")
	log.Debug("this is e test")

	time.Sleep(time.Second * 10)

}
