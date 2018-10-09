package clog

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/toolkits/file"
)

const (
	L_BLACK  = `\e[1;30m`
	RED      = `\e[0;31m`
	L_RED    = `\e[1;31m`
	GREEN    = `\e[0;32m`
	L_GREEN  = `\e[1;32m`
	BROWN    = `\e[0;33m`
	YELLOW   = `\e[1;33m`
	BLUE     = `\e[0;34m`
	L_BLUE   = `\e[1;34m`
	PURPLE   = `\e[0;35m`
	L_PURPLE = `\e[1;35m`
	CYAN     = `\e[0;36m`
	L_CYAN   = `\e[1;36m`
	GRAY     = `\e[0;37m`
	WHITE    = `\e[1;37m`
)

var (
	timeZone *time.Location
	err      error
	ff       *os.File
)

func logout(l *Logger) (err error) {
	defer func() {
		err := recover()
		fmt.Fprintf(os.Stdout, "\033[0;31m[%s] [%s]  %s\033[0m\n", time.Now().Format("2006-01-02 15:04:05"), "Error", err)
		close(l.Out)
		ff.Close()
		os.Exit(2)

	}()

	f := filepath.Join(l.LogPath, l.LogFileName)
	ff, err = os.OpenFile(f, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0744)
	if err != nil {
		fmt.Fprintf(os.Stdout, "\033[0;31m[%s] [%s]  %s\033[0m\n", time.Now().Format("2006-01-02 15:04:05"), "Error", err)
		// continue
		ff.Close()
	}

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
			os.Exit(2)
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

		if l.EnableFileLog {
			if file.IsExist(l.LogPath) {
				if l.LogLevel >= msg.LogLevel {
					// 避免重复打开关闭文件描述符,提高性能
					// fmt.Println(msg.Msg)
					// file.WriteBytes(filepath.Join(l.LogPath, l.LogFileName),[]byte(fmt.Sprint(msg.Msg)))
					// f := filepath.Join(l.LogPath, l.LogFileName)
					// ff, err := os.OpenFile(f, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0744)
					// if err != nil {
					// 	fmt.Fprintf(os.Stdout, "\033[0;31m[%s] [%s]  %s\033[0m\n", msg.Time.In(timeZone).Format("2006-01-02 15:04:05"), "Error", err)
					// 	continue
					// }
					fw := bufio.NewWriter(ff)
					fw.WriteString(fmt.Sprintln(msg.Msg))
					// fmt.Println(n, "#", err)
					fw.Flush()
					// ff.Close()
					// fmt.Println("日志文件已关闭")
				}
			} else {
				fmt.Fprintf(os.Stdout, "\033[0;31m[%s] [%s]  %s\033[0m\n", msg.Time.In(timeZone).Format("2006-01-02 15:04:05"), "Error", "Log file path not exist")

			}
		}

	}
	return
}
