package stack

type Stack struct {
	arr []int
}

func (s *Stack) Push(a int) {
	s.arr = append(s.arr, a)
}
func (s *Stack) Pop() int {
	last := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return last
}
func (s *Stack) Top() int {
	return s.arr[len(s.arr)-1]
}
func (s *Stack) Len() int {
	return len(s.arr)
}

func main() {
	//s := Stack{}
	//s.Push(1)
	//s.Push(2)
	//s.Push(3)
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
}
