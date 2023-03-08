package main

import (
	"net/http"
)

func main() {
	n := 100
	//wg := sync.WaitGroup{}
	//wg.Add(n)
	//for i := 0; i < n; i++ {
	//	go func(i int) {
	//		http.Get("http://www.baidu.com")
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	for i := 0; i < n; i++ {
		http.Get("http://www.baidu.com")
	}
}
