package main

type MyQueue struct {
	arr []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyQueue) Pop() int {
	t := this.arr[0]
	this.arr = this.arr[1:]
	return t
}

func (this *MyQueue) Peek() int {
	return this.arr[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}

func main() {

}
