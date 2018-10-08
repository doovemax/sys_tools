package main

import (
	"fmt"
	"github.com/doovemax/sys_tools/clog"
	"time"
)

func main() {
	fmt.Println(clog.Clog)
	clog.Clog.Fatal("this is a test")

	time.Sleep(1)

}
