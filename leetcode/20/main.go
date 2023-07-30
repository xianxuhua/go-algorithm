package main

import "fmt"

type stack struct {
	arr []int
}

func (s *stack) push(i int) {
	s.arr = append(s.arr, i)
}

func (s *stack) pop() int {
	if len(s.arr) == 0 {
		return -1
	}
	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return r
}

func (s *stack) top() int {
	if len(s.arr) == 0 {
		return -1
	}
	return s.arr[len(s.arr)-1]
}

func isValid(s string) bool {
	sta := stack{}
	pushedCount := 0
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			sta.push(int(v))
			pushedCount++
		} else {
			t := sta.top()
			if t == '(' && v == ')' || t == '[' && v == ']' || t == '{' && v == '}' {
				sta.pop()
			}
		}
	}
	if pushedCount*2 == len(s) && len(sta.arr) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("]"))
	fmt.Println(isValid("(])"))

}
