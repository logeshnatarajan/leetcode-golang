package main

import (
	"fmt"
	"sort"
)

func main() {
	input := "qrssaqqzibbssraib"
	m1, m2 := map[string]int{}, map[int]string{}
	for _, v := range input {
		m1[string(v)]++
	}
	var res []int
	for k, v := range m1 {
		res = append(res, v)
		m2[v] = k
	}
	sort.Ints(res)
	result := ""
	for i := len(res) - 1; i >= 0; i-- {
		temp := ""
		for k, v := range m1 {
			if v == res[i] {
				temp += k
				delete(m1, k)
			}
		}
		result += SortString(temp)
	}
	fmt.Println(result)
}
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
