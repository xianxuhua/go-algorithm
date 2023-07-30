package main

import (
	"context"
	"fmt"
	"time"
)

type key struct{}

func main() {
	c := context.Background()
	c = context.WithValue(c, key{}, "value")
	c, cancel := context.WithTimeout(c, time.Second*3)
	defer cancel()
	go mainTask(c)

	// 手动cancel
	var cmd string
	for {
		fmt.Scan(&cmd)
		if cmd == "c" {
			cancel()
		}
	}

}

func mainTask(c context.Context) {
	// 后台任务cancel方法
	go func() {
		c2, cancel2 := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel2()
		// 后台任务
		smallTask(c2, "go task", time.Second*5)
	}()

	smallTask(c, "small process", time.Second*2)
}

func smallTask(c context.Context, name string, timeout time.Duration) {
	select {
	case <-c.Done():
		fmt.Println(name, "timeout")
	case <-time.After(timeout):
		fmt.Println(name, "done")
	}
}
