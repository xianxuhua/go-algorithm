package main

import (
	"fmt"
	"time"
)

func main() {
	var a [1000]int
	// goroutine切换点：io/select,channel,等待锁,go1.14非io也能切换
	for i := 0; i < 1000; i++ {
		go func(i int) {
			// go run -race main.go 检测数据访问冲突
			a[i]++
			fmt.Println("hello")
		}(i)
	}
	time.Sleep(time.Minute)
}
