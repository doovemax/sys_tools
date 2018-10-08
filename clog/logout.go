package clog

import (
	"fmt"
	"os"
	"time"
)

var (
	timeZone *time.Location
	err      error
)

func logout(l *logger) (err error) {
	timeZone, err = time.LoadLocation(l.Timezone)
	if err != nil {
		fmt.Fprintf(os.Stdout, "\033[7;31m[%s] TimeZone ERROR default \"Local\"\033[0m\n", time.Now().Format("2006-01-02 15:04:05"))
		timeZone, _ = time.LoadLocation("Local")
	}

	for msg := range l.Out {
		//默认后台全部输出
		//文件输出
		switch msg.LogLevel {
		case FatalLevel:
			fmt.Fprintf(os.Stdout, "\033[7;31m[%s] [%s]  %s\033[0m\n", msg.Time.In(timeZone).Format("2006-01-02 15:04:05"), "Fatal", msg.Msg)
		case ErrorLevel:
			fmt.Fprintf(os.Stdout, "\033[0;31m[%s] [%s]  %s\033[0m\n", msg.Time.In(timeZone).Format("2006-01-02 15:04:05"), "Error", msg.Msg)
		case WarnLevel:
			fmt.Fprintf(os.Stdout, "\033[0;33m[%s] [%s]  %s\033[0m\n", msg.Time.In(timeZone).Format("2006-01-02 15:04:05"), "Warn", msg.Msg)
		case InfoLevel:
			fmt.Fprintf(os.Stdout, "\033[0;32m[%s] [%s]  %s\033[0m\n", msg.Time.In(timeZone).Format("2006-01-02 15:04:05"), "Info", msg.Msg)
		case DebugLevel:
			fmt.Fprintf(os.Stdout, "\033[0;35m[%s] [%s]  %s\033[0m\n", msg.Time.In(timeZone).Format("2006-01-02 15:04:05"), "Debug", msg.Msg)
		default:
			fmt.Println(msg, l)
		}

	}
	return
}