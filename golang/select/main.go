package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))
			c <- i
			i++
		}
	}()
	return c
}

func main() {
	c1, c2 := generator(), generator()
	for {
		select {
		case <-c1:
			fmt.Println("received from c1")
		case <-c2:
			fmt.Println("received from c2")
		case <-time.After(time.Millisecond * 1000):
			fmt.Println("timeout")
		}
	}
}
