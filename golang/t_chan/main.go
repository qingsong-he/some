package main

import (
	. "github.com/qingsong-he/ce"
	"time"
)

func Case1() {
	ch1 := make(chan int)
	exitCh := make(chan int)

	go func() {
		v, ok := <-ch1
		Print(v, ok)
		exitCh <- 0
	}()

	close(ch1)
	<-exitCh
	v, ok := <-ch1
	Print(v, ok)
}

func Case2() {
	ch1 := make(chan int)
	exitCh := make(chan int)

	go func() {
		ch1 <- 1 // panic
		exitCh <- 0
	}()

	close(ch1)
	<-exitCh
	v, ok := <-ch1
	Print(v, ok)
}

func Case3() {
	ch1 := make(chan int)
	select {
	case v, ok := <-ch1:
		Print(v, ok)
	default:
		Print(".")
	}
}

func Case4() {
	var ch1 chan int
	v, ok := <-ch1 // panic
	Print(v, ok)
}

func Case5() {
	var ch1 chan int
	ch1 <- 0 // panic
}

func Case6() {
	var ch1 chan int
	close(ch1) // panic
}

func Case7() {
	ch1 := make(chan int)
	close(ch1)
	v, ok := <-ch1
	Print(v, ok)
}

func Case8() {
	ch1 := make(chan int)
	close(ch1)
	ch1 <- 0 // panic
}

func Case9() {
	ch1 := make(chan int)
	close(ch1)
	close(ch1) // panic
}

func Case10() {
	ch1 := make(chan int, 0xff)
	ch1 <- 1
	ch1 <- 2
	close(ch1)

	v, ok := <-ch1
	Print(v, ok)

	v, ok = <-ch1
	Print(v, ok)

	v, ok = <-ch1
	Print(v, ok)
}

func Case11() {
	ch1 := make(chan int, 0)
	go func() {
		<-ch1
		Print("return 1")
	}()
	go func() {
		<-ch1
		Print("return 2")
	}()

	time.Sleep(1 * time.Second)
	close(ch1)
	time.Sleep(1 * time.Second)
}

func main() {
	Case1()
}
