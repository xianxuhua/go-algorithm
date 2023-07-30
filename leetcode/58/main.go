package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	split := strings.Split(strings.TrimSpace(s), " ")
	return len(split[len(split)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
