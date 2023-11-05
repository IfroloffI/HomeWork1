package main

import (
	"fmt"
	"sort"
)

type sortInt []int

func (s sortInt) Len() int {
	return len(s)
}

func (s sortInt) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s sortInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	arr := []int{a, b, c}
	sort.Sort(sortInt(arr))
	if arr[0] < arr[1]+arr[2] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
