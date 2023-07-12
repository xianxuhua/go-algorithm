package main

import (
	"fmt"
)

func main() {
	workersDemo()
}

type Worker struct {
	c     chan int
	close chan bool
}

// consumer
func doWorker(id int, worker Worker) {
	// wait data
	//for {
	//	fmt.Printf("worker %d received %c\n", id, <-c)
	//}
	for i := range worker.c {
		fmt.Printf("worker %d received %c\n", id, i)
		worker.close <- true
	}
}

// return: write only channel
func createWorker(id int) Worker {
	//c := make(chan int)
	c := Worker{close: make(chan bool), c: make(chan int)}
	go doWorker(id, c)
	return c
}

func workersDemo() {
	workers := [10]Worker{}
	// create workers
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	// send to channels
	for i := 0; i < 10; i++ {
		workers[i].c <- i + 'a'
	}
	// wait all worker close
	for _, worker := range workers {
		<-worker.close
	}
	for i := 0; i < 10; i++ {
		workers[i].c <- i + 'A'
	}
	// wait all worker close
	for _, worker := range workers {
		<-worker.close
	}
}
