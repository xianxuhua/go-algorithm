package main

import (
	"fmt"
	"strings"
)

type stack struct {
	arr []string
}

func (s *stack) push(i string) {
	s.arr = append(s.arr, i)
}

func (s *stack) pop() string {
	if len(s.arr) == 0 {
		return ""
	}
	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return r
}

func (s *stack) top() string {
	if len(s.arr) == 0 {
		return ""
	}
	return s.arr[len(s.arr)-1]
}

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	s := stack{}
	for i := 0; i < len(split); i++ {
		if split[i] == ".." && len(s.arr) > 0 {
			s.pop()
		} else if split[i] != "" && split[i] != "." && split[i] != ".." {
			s.push(split[i])
		}
	}
	return "/" + strings.Join(s.arr, "/")
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
