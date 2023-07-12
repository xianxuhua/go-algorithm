package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) incr() {
	fmt.Println("safe incr")
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {
	a := atomicInt{value: 0, lock: sync.Mutex{}}
	// go run -race main.go
	go func() {
		a.incr()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
