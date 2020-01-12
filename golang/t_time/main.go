package main

import (
	. "github.com/qingsong-he/ce"
	"time"
)

func Case1() {
	ticker1 := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			// 'ok' is always true
			case time1, ok := <-ticker1.C:
				Print(time1, ok)
				time.Sleep(2 * time.Second) // ignore some alarm from 'ticker1'
			}
		}
	}()

	time.Sleep(6 * time.Second)
	ticker1.Stop()
}

func Case2() {
	timer1 := time.NewTimer(2 * time.Second)
	go func() {
		for {
			select {
			case time1, ok := <-timer1.C:
				Print(time1, ok)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	Print(timer1.Stop()) // false (expired)
}

func Case3() {
	timer1 := time.NewTimer(2 * time.Second)
	go func() {
		for {
			select {
			case time1, ok := <-timer1.C:
				Print(time1, ok)
			}
		}
	}()
	Print(timer1.Stop()) // true (not expired, not stopped)
	Print(timer1.Stop()) // false (stopped)
}

func Case4() {
	timer1 := time.NewTimer(2 * time.Second)
	go func() {
		for {
			select {
			case time1, ok := <-timer1.C:
				Print(time1, ok)
			}
		}
	}()
	Print(timer1.Reset(4 * time.Second)) // true (not expired, not stopped)
	time.Sleep(5 * time.Second)
	Print(timer1.Reset(4 * time.Second)) // false (expired)
}

func Case5() {
	utcTime := time.Date(2018, 1, 1, 1, 1, 1, 0, time.UTC)
	localTime := time.Date(2018, 1, 1, 1, 1, 1, 0, time.Local)
	Print(utcTime, utcTime.Unix())     // 2018-01-01 01:01:01 +0000 UTC 1514768461
	Print(localTime, localTime.Unix()) // 2018-01-01 01:01:01 +0700 +07 1514743261

	chinaZone := time.FixedZone("+8", 8*60*60)
	Print(utcTime.In(chinaZone), utcTime.In(chinaZone).Unix())     // 2018-01-01 09:01:01 +0800 +8 1514768461
	Print(localTime.In(chinaZone), localTime.In(chinaZone).Unix()) // 2018-01-01 02:01:01 +0800 +8 1514743261
}

func main() {
	Case5()
}
