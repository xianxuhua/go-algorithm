package main

import "fmt"

type MyStack struct {
	arr []int
}

func Constructor() MyStack {
	return MyStack{arr: []int{}}
}

func (this *MyStack) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyStack) Pop() int {
	t := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return t
}

func (this *MyStack) Top() int {
	t := this.arr[len(this.arr)-1]
	return t
}

func (this *MyStack) Empty() bool {
	return len(this.arr) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
	fmt.Println(obj.Empty())
}
