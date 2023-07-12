package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createConsumer(id int) chan int {
	c := make(chan int)
	go func(id int) {
		for {
			fmt.Printf("consumer %d received %c\n", id, <-c)
		}
	}(id)
	return c
}

func createProducer(id int) chan int {
	c := make(chan int)
	go func(id int) {
		i := 0
		for {
			// spend some time to product data
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			fmt.Printf("producer %d generate %c\n", id, 'a'+i)
			c <- 'a' + i
			i++
		}
	}(id)
	return c
}

func main() {
	c1, c2 := createProducer(0), createProducer(1)
	c3 := createConsumer(0)
	for {
		n := 0
		select {
		case n = <-c1:
			c3 <- n
			//fmt.Printf("received from c1: %c\n", n)
		case n = <-c2:
			c3 <- n
			//fmt.Printf("received from c2: %c\n", n)
		}
	}
}
