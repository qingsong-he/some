package main

import (
	"github.com/qingsong-he/ce"
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
			ce.Print(err)
			runtime.Gosched()
			continue
		}

		recvBuf := make([]byte, 1)
		go func() {
			defer func() {
				if errByPanic := recover(); errByPanic != nil {
					if _, ok := ce.IsFromCe(errByPanic); !ok {
						ce.Print(errByPanic)
					}
				}
			}()
			for {
				err := conn.SetReadDeadline(time.Now().Add(time.Second * 60))
				ce.CheckError(err)

				_, err = io.ReadFull(conn, recvBuf)
				ce.CheckError(err)

				_, err = conn.Write(recvBuf)
				ce.CheckError(err)
			}
		}()
	}

}
