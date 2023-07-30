package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[int(s[i])]++
	}
	indexes := []int{}
	// record repeat indexes
	for i := 0; i < len(s); i++ {
		if m[int(s[i])] != 0 {
			indexes = append(indexes, i)
		}
	}
	temp := make(map[int]int)
	for _, v := range indexes {
		if _, ok := temp[int(t[v])]; !ok {
			temp[int(t[v])]++
		}
	}
	return len(temp) == 1
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("a", "a"))
	fmt.Println(isIsomorphic("ab", "ac"))
}
