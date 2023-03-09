package main

import (
	"fmt"
	"math/rand"
)

type Set struct {
	arr []int
}

func (s *Set) Add(e int) {
	contains := false
	for _, v := range s.arr {
		if v == e {
			contains = true
		}
	}
	if !contains {
		s.arr = append(s.arr, e)
	}
}

func (s *Set) Remove(e int) {
	for i := 0; i < len(s.arr); i++ {
		if s.arr[i] == e {
			s.arr = append(s.arr[:i], s.arr[i+1:]...)
		}
	}
}

func main() {
	s := Set{}
	for i := 0; i < 10; i++ {
		s.Add(rand.Intn(10))
	}
	fmt.Println(s.arr)
	s.Remove(1)
	fmt.Println(s.arr)
}
