package main

import (
	. "github.com/qingsong-he/ce"
	"sync"
	"time"
)

var sm sync.Map

func func1(user string) {
	_, ok := sm.LoadOrStore(user, struct{}{})
	if ok {
		Print(user, "busy")
		return
	}

	defer func() {
		sm.Delete(user)
	}()

	Print(user, "free")
}

func main() {
	for i := 0; i < 0xf; i++ {
		go func1("user1")
	}
	time.Sleep(10 * time.Second)
}
