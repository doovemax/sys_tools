package main

import (
	"fmt"
	"time"

	"github.com/doovemax/sys_tools/clog"
)

func main() {
	fmt.Println(clog.Clog)
	clog.Clog.Fatal("this is a test")
	clog.Clog.Error("this is a test")
	clog.Clog.Warn("this is a test")
	clog.Clog.Info("this is a test")
	clog.Clog.Debug("this is a test")
	//clog.Clog.Panic("this is a test")

	time.Sleep(time.Second * 10)
	fmt.Println()

}
