package main

import (
	"expvar"
	"github.com/qingsong-he/ce"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func init() {
	ce.Print(os.Args[0])
}

type Obj1 struct {
	sync.RWMutex
	m map[uint64]uint64
	s []int
}

func (o *Obj1) clean() {
	for k := range o.m {
		delete(o.m, k)
	}
	o.s = nil
}

var m1 = &sync.Map{}
var p1 = &sync.Pool{
	New: func() interface{} {
		return &Obj1{
			m: make(map[uint64]uint64),
		}
	},
}

func Case1(n uint64) {
	for i := uint64(0); i < n; i++ {
		go func(i uint64) {
			expvarByI := expvar.NewString(strconv.Itoa(int(i)))
			for {
				o1 := p1.Get().(*Obj1)
				o1.clean()

				if actO1, ok := m1.LoadOrStore(i, o1); ok {
					p1.Put(o1)
					o1 = nil
					o1 = actO1.(*Obj1)
				} else {
					print(i, ",")
				}

				func() {
					o1.Lock()
					defer o1.Unlock()
					o1.m[i]++
					expvarByI.Set(strconv.FormatUint(o1.m[i], 10))
				}()

				time.Sleep(time.Duration(rand.Uint64()%1000) * time.Millisecond)
			}
		}(i)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	Case1(10000)
	http.ListenAndServe(":3000", nil)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	ce.Print(<-ch)
}
