package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	minLenStr := strs[0]
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
			minLenStr = str
		}
	}

	// to f, f, f, l, l, l, o, o, i, ...
	l := []int{}
	for i := 0; i < minLen; i++ {
		for _, str := range strs {
			l = append(l, int(str[i]))
		}
	}

	maxPreIndex := 0
	for i := 0; i < len(l)/len(strs); i++ {
		// [i...i+len(strs)-1] equal ?
		m := make(map[int]int)
		for j := 0; j < len(strs); j++ {
			m[l[i*len(strs)+j]]++
		}
		if len(m) == 1 {
			maxPreIndex = i + 1
		} else {
			break
		}
	}

	return minLenStr[:maxPreIndex]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}
