package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	return find(text1, text2, len(text1)-1, len(text2)-1)
}

func find(text1, text2 string, index1, index2 int) int {
	if index1 < 0 || index2 < 0 {
		return 0
	}

	res := 0
	if text1[index1] == text2[index2] {
		res = 1 + find(text1, text2, index1-1, index2-1)
	} else {
		res = max(find(text1, text2, index1-1, index2), find(text1, text2, index1, index2-1))
	}
	return res
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	cache := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		cache[i] = make([]int, len(text2)+1)
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				cache[i+1][j+1] = 1 + cache[i][j]
			} else {
				cache[i+1][j+1] = max(cache[i][j+1], cache[i+1][j])
			}
		}
	}

	return cache[len(text1)][len(text2)]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("ezupkr", "ubmrapg"))
	fmt.Println(longestCommonSubsequence2("abcde", "ace"))
	fmt.Println(longestCommonSubsequence2("ezupkr", "ubmrapg"))
}
