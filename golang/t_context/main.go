package main

import (
	"context"
	. "github.com/qingsong-he/ce"
	"time"
)

func Case1() {
	c1 := context.Background()

	c2, cancelFunc2 := context.WithCancel(c1)
	Print(cancelFunc2)

	c3, cancelFunc3 := context.WithCancel(c2)
	Print(cancelFunc3)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			Print("done 1")
			Print(ctx.Deadline())
			Print(ctx.Err())
		}
	}(c2)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			Print("done 2")
			Print(ctx.Deadline())
			Print(ctx.Err())
		}
	}(c3)

	time.Sleep(3 * time.Second)
	cancelFunc2()
	time.Sleep(3 * time.Second)
}

func Case2() {
	c1 := context.Background()

	c2, cancelFunc2 := context.WithDeadline(c1, time.Now().Add(6*time.Second))
	Print(cancelFunc2)

	c3, cancelFunc3 := context.WithDeadline(c2, time.Now().Add(6*time.Second))
	Print(cancelFunc3)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			Print("done 1")
			Print(ctx.Deadline())
			Print(ctx.Err())
		}
	}(c2)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			Print("done 2")
			Print(ctx.Deadline())
			Print(ctx.Err())
		}
	}(c3)

	time.Sleep(3 * time.Second)
	cancelFunc2()
	time.Sleep(3 * time.Second)
}

func main() {
	Case1()
}
