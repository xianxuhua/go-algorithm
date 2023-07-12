package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))
			c <- fmt.Sprintf("msg: generator %s: %c", name, 'a'+i)
			i++
		}
	}()
	return c
}

// aggregate channels to one channel
func fanIn(c1, c2 chan string) chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-c1
		}
	}()
	go func() {
		for {
			out <- <-c2
		}
	}()
	return out
}

func fanIn2(chs ...chan string) chan string {
	out := make(chan string)
	for _, c := range chs {
		c := c // golang bug
		go func() {
			for {
				out <- <-c
			}
		}()
	}
	return out
}

func fanInWithSelect(c1, c2 chan string) chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case n := <-c1:
				out <- n
			case n := <-c2:
				out <- n
			}
		}
	}()
	return out
}

func main() {
	c1 := generator("c1")
	c2 := generator("c2")
	c := fanIn2(c1, c2)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(time.Millisecond * 1000):
			fmt.Println("no received from c")
		}
	}

}
