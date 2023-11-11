package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	first := arr[size-1]
	for i := size - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = first
	fmt.Println(arr)
}
