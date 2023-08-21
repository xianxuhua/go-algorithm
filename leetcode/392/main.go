package main

import "fmt"

func isSubsequence(s string, t string) bool {
	si, ti := 0, 0
	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
			ti++
		} else {
			ti++
		}
	}
	return si == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
