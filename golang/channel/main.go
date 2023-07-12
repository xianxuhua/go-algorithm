package main

import (
	"fmt"
	"time"
)

func main() {
	closableChannelDemo()
}

func errorDemo() {
	c := make(chan int)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func goodDemo() {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func workerDemo() {
	c := make(chan<- int)
	for i := 0; i < 10; i++ {
		c = createWorker(i)
	}
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Millisecond)
}

// consumer
func worker(id int, c chan int) {
	// wait data
	//for {
	//	fmt.Printf("worker %d received %c\n", id, <-c)
	//}
	for i := range c {
		fmt.Printf("worker %d received %c\n", id, i)
	}
}

// return: write only channel
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func workersDemo() {
	channels := [10]chan<- int{}
	// create workers
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	// send to channels
	for i := 0; i < 10; i++ {
		channels[i] <- i + 'a'
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i + 'A'
	}
	time.Sleep(time.Millisecond)
}

func buffDemo() {
	// buffer
	c := make(chan int, 1)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	time.Sleep(time.Millisecond)
}

func closableChannelDemo() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	close(c) // channel close, goroutine for range
	time.Sleep(time.Millisecond)
}
