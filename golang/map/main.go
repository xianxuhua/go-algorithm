package main

import (
	"fmt"
	"sync"
	"time"
)

func nilMap() {
	var m map[string]int
	fmt.Println(m["1"]) // 能取值
	m["1"] = 1          //panic: assignment to entry in nil map
}

// map对象不是线程安全的

type Map struct {
	m    map[string]int
	lock sync.Mutex
}

func (m *Map) put(key string, value int) {
	m.lock.Lock()
	m.m[key] = value
	m.lock.Unlock()
}

func (m *Map) get(key string) int {
	m.lock.Lock()
	defer m.lock.Unlock()

	return m.m[key]
}

func forRange() {
	// 循环时也无序
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 2
	// delete不会释放内存
	delete(m, "a")

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func concurrent() {
	m := Map{m: make(map[string]int), lock: sync.Mutex{}}
	go func() {
		m.put("a", 1)
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(m.get("a"))
}

func key() {
	//可比较的类型都可以作为map key
	//不能作为map key 的类型包括：
	//slices
	//maps
	//functions
	m := make(map[chan int]int)
	c := make(chan int)
	m[c] = 1
	fmt.Println(m)
}

func main() {
	//nilMap()

}
