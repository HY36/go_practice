package main

import (
	"fmt"
	"sort"
)

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

func shortestToChar(S string, C byte) []int {
	var result []int
	front, last := 0, 0
	for index, value := range S {
		if byte(value) == C {
			last = index
			for i := 0; i <= (last-front)/2; i++ {
				result = append(result, i)
			}
			front = index
		}
	}
	return result
}

func sortedSquares(A []int) []int {
	for indx, value := range A {
		A[indx] = value * value
	}
	sort.Ints(A)
	return A
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	result := 0
	N := len(nums)
	for i := 0; i < N; i += 2 {
		result += nums[i]
	}
	return result
}

func hasAlternatingBits(n int) bool {

	return false
}

func repeatedNTimes(A []int) int {
	container := make(map[int]int, len(A)/2+1)
	for _, value := range A {
		if _, ok := container[value]; !ok {
			container[value] = 1
		} else {
			return value
		}
	}
	return 0
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	result := make([]int, len(A))
	evenMaps := make(map[int]int)
	evenValue := 0
	for index, value := range A {
		if value%2 == 0 {
			evenValue += value
			evenMaps[index] = value
		}
	}
	for ind, value := range queries {
		index := value[1]
		val := value[0]
		if (A[index]+val)%2 == 0 {
			if _, ok := evenMaps[index]; ok {
				evenValue += val
			} else {
				evenValue += A[index] + val
				evenMaps[index] = 0
			}
		} else {
			if _, ok := evenMaps[index]; ok {
				evenValue -= A[index]
				delete(evenMaps, index)
			}
		}
		A[index] += val
		result[ind] = evenValue
	}
	return result
}

func smallestRangeI(A []int, K int) int {
	maxValue, minValue := A[0], A[0]
	for _, value := range A {
		if value > maxValue {
			maxValue = value
		}
		if value < minValue {
			minValue = value
		}
	}
	dValue := maxValue - minValue
	if dValue <= 2*K {
		return 0
	}
	return dValue - 2*K
}

func main() {
	A := []int{2,7,2}
	fmt.Println(smallestRangeI(A, 1))
}
