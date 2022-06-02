package main

import (
	"fmt"
	"sort"
	"strings"
)

//Input - “readable content of a page when looking at its layout”

//Output - “readable content looking layout page when its of at a”
func main() {
	input := strings.Fields("readable content of a page when looking at its layout")
	m1, m2 := map[string]int{}, map[int]string{}
	for _, v := range input {
		m1[v] = len(v)
	}
	var res []int
	for k, v := range m1 {
		res = append(res, v)
		m2[v] = k
	}
	sort.Ints(res)
	result := []string{}
	for i := len(res) - 1; i >= 0; i-- {
		temp := []string{}
		for k, v := range m1 {
			if v == res[i] {
				temp = append(temp, k)
				delete(m1, k)
			}
		}
		sort.Strings(temp)
		for _, val := range temp {
			result = append(result, val)
		}
	}
	fmt.Println(strings.Join(result, " "))
}
