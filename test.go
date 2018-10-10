package main

import (
	"runtime"
	"sync"
	"time"

	"github.com/doovemax/sys_tools/clog"
)

func main() {
	runtime.GOMAXPROCS(1)
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
		Sizecron:      0,
		Timecron:      "*/20 * * * * *",
	}
	log.Run()
	for {
		// log.Fatal("this is a test")
		log.Error("this is b test")
		log.Warn("this is c test")
		log.Info("this is d test")
		log.Debug("this is e test")
		// size, _ := file.FileSize("/tmp/clog.log")
		// if size >= 2048 {
		time.Sleep(time.Second * 10)
		// }
	}

}
