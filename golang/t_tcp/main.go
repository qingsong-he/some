package main

import (
	"github.com/qingsong-he/ce"
	"go.uber.org/zap"
	"io"
	"net"
	"os"
	"runtime"
	"time"
)

func init() {
	ce.Print(os.Args[0])
}

func main() {
	l, err := net.Listen("tcp", ":9999")
	ce.CheckError(err)
	for {
		conn, err := l.Accept()
		if err != nil {
			ce.Error("", zap.Error(err))
			runtime.Gosched()
			continue
		}

		recvBuf := make([]byte, 1)
		go func() {
			defer func() {
				if errByPanic := recover(); errByPanic != nil {
					if !ce.IsFromMe(errByPanic) {
						ce.Error("", zap.Any("errByPanic", errByPanic))
					}
				}
			}()
			for {
				conn.SetReadDeadline(time.Now().Add(time.Second * 60))
				_, err := io.ReadFull(conn, recvBuf)
				ce.CheckError(err)
				conn.Write(recvBuf)
			}
		}()
	}

}
