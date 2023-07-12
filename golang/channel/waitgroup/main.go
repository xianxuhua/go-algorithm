package main

import (
	"fmt"
	"sync"
)

func main() {
	workersDemo()
}

type Worker struct {
	c    chan int
	done func()
}

// consumer
func doWorker(id int, worker Worker) {
	// wait data
	//for {
	//	fmt.Printf("worker %d received %c\n", id, <-c)
	//}
	for i := range worker.c {
		fmt.Printf("worker %d received %c\n", id, i)
		worker.done()
	}
}

// return: write only channel
func createWorker(id int, wg *sync.WaitGroup) Worker {
	//c := make(chan int)
	worker := Worker{done: func() {
		wg.Done()
	}, c: make(chan int)}
	go doWorker(id, worker)
	return worker
}

func workersDemo() {
	workers := [10]Worker{}
	wg := sync.WaitGroup{}
	// create workers
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)
	// send to channels
	for i := 0; i < 10; i++ {
		workers[i].c <- i + 'a'
	}
	for i := 0; i < 10; i++ {
		workers[i].c <- i + 'A'
	}
	// wait all worker close
	wg.Wait()
}
