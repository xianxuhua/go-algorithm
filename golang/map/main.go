package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrent() {
	//map对象不是线程安全的
	m := make(map[int]int)
	lock := sync.Mutex{}
	for i := 0; i < 50; i++ {
		i := i
		go func() {
			//fatal error: concurrent map writes
			lock.Lock()
			m[i]++
			lock.Unlock()
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(m)
}

func nilMap() {
	var m map[string]int
	fmt.Println(m["1"]) // 能取值
	m["1"] = 1          //panic: assignment to entry in nil map
}

func main() {
	//concurrent()
	nilMap()
}
