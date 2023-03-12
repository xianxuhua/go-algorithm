package main

import "fmt"

type Item struct {
	count int
	index int
}

func minWindow(s string, t string) string {
	m1, m2 := make(map[rune]int), make(map[rune]Item)
	i, j := 0, 0
	res := ""
	minLen := len(s)
	for _, v := range t {
		m2[v] = Item{
			count: m2[v].count + 1,
			index: 0,
		}
	}
	for j < len(s) {
		if m2[rune(s[j])].count > 0 && m1[rune(s[j])] == 0 {
			m1[rune(s[j])]++
			m2[rune(s[j])] = Item{
				count: m2[rune(s[j])].count + 1,
				index: j,
			}
			j++
		} else if m2[rune(s[j])].count == 0 {
			j++
		} else if m2[rune(s[j])].count > 0 && m1[rune(s[j])] > 0 {
			// find max index
			minIndex := m2[0].index
			fmt.Println(j)
			for k := range m2 {
				fmt.Printf("%c, count %d, index %d, ", k, m2[k].count, m2[k].index)
			}
			fmt.Println()
			for k := range m2 {
				if m2[k].index > minIndex {
					minIndex = m2[k].index
				}
			}
			for m := i; m < minIndex; m++ {
				delete(m1, rune(s[m]))
			}
			i = minIndex
		}

		contain := true
		for k := range m2 {
			if m1[k] != m2[k].count {
				contain = false
			}
		}
		if contain && j-i+1 < minLen {
			res = s[i : j+1]
		}
	}
	return res
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
