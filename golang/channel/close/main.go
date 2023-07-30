package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		fmt.Println(<-c)
	}()
	c <- 1
	close(c)
	//关闭的管道仍然可以读数据
	time.Sleep(time.Millisecond)
}
