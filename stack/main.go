package stack

type Stack struct {
	arr []string
}

func (s *Stack) Push(a string) {
	s.arr = append(s.arr, a)
}
func (s *Stack) Pop() string {
	last := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return last
}
func (s *Stack) Top() string {
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
