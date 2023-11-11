package main

import "fmt"

func createMatrix(arr *[][]int, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			(*arr)[i] = append((*arr)[i], 0)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([][]int, n)
	createMatrix(&arr, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] != arr[n-i-1][n-j-1] {
				println("no")
				return
			}
		}
	}
	println("yes")
}
