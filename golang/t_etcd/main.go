package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/qingsong-he/ce"
	"os"
	"time"
)

func init() {
	ce.Print(os.Args[0])
}

func concurrencyTest(pfx string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	ce.CheckError(err)
	defer cli.Close()

	se, err := concurrency.NewSession(cli)
	ce.CheckError(err)
	defer se.Close()

	mu := concurrency.NewMutex(se, pfx)
	err = mu.Lock(context.TODO())
	ce.CheckError(err)
	ce.Print("get lock")
	time.Sleep(5 * time.Second)

	err = mu.Unlock(context.TODO())
	ce.CheckError(err)
	ce.Print("set unlock")
}

func main() {
	concurrencyTest(os.Args[1])
}
