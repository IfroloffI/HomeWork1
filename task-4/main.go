package main

import "fmt"

func createMatrix(arr *[][]int, size *int) {
	for i := 0; i < *size; i++ {
		for j := 0; j < *size; j++ {
			(*arr)[i] = append((*arr)[i], 0)
		}
	}
}

func transpose(arr *[][]int, size *int) [][]int {
	trArr := make([][]int, *size)
	createMatrix(&trArr, size)
	for i := 0; i < *size; i++ {
		for j := 0; j < *size; j++ {
			trArr[i][j] = (*arr)[j][i]
		}
	}
	return trArr
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([][]int, n)
	createMatrix(&arr, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	trArr := transpose(&arr, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] != trArr[i][j] {
				println("no")
				return
			}
		}
	}
	println("yes")
}
