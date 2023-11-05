package main

import (
	"fmt"
)

func main() {
	var len, number int
	fmt.Scan(&len)
	var arr []int
	for i := 0; i < len; i++ {
		fmt.Scan(&number)
		arr = append(arr, number)
	}
	first := arr[len-1]
	for i := len - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = first
	fmt.Println(arr)
}
