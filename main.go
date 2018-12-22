package main

import "fmt"

func transpose(A [][]int) [][]int {
	line, column := len(A), len(A[0])
	var B [][]int
	for i := 0; i < column; i++ {
		var tmp []int
		for j := 0; j < line; j++ {
			tmp = append(tmp, A[j][i])
		}
		B = append(B, tmp)
	}
	return B
}

func main() {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	B := transpose(A)
	fmt.Println(B)
}
