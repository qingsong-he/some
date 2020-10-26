package main

import (
	"github.com/qingsong-he/ce"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func init() {
	ce.SetOutput(
		io.MultiWriter(
			&lumberjack.Logger{
				Filename:   "/tmp/foobar.log",
				MaxSize:    100, // Size(megabytes) of each log file
				MaxAge:     28,  // Days
				MaxBackups: 3,   // Maximum number of log files
			},
			os.Stderr,
		),
	)
	ce.Print(os.Args[0])
}

func main() {
	for i := 0; i < 1; i++ {
		ce.Print(os.Args[0])
	}
	defer ce.Sync()
}
