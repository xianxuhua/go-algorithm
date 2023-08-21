package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	m := []string{
		" ",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var res []string
	findCombine(digits, 0, "", m, func(tmp string) {
		res = append(res, tmp)
	})
	return res
}

func findCombine(digits string, index int, tmp string, m []string, opt func(tmp string)) {
	if index == len(digits) {
		opt(tmp)
		return
	}
	cur := digits[index]
	letters := m[cur-'0']
	for _, letter := range letters {
		findCombine(digits, index+1, tmp+fmt.Sprintf("%c", letter), m, opt)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
