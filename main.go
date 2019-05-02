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

func diStringMatch(S string) []int {
	result := make([]int, len(S)+1)
	leftIndex, rightIndex := 0, len(S)
	index := 0
	for _, value := range S {
		if value == 'D' {
			result[index] = rightIndex
			rightIndex -= 1
		} else {
			result[index] = leftIndex
			leftIndex += 1
		}
		index += 1
	}
	result[index] = leftIndex
	return result
}

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	nums2Map := make(map[int]int)
	result := make([]int, 0)
	nums1Length := len(nums1)
	nums2Length := len(nums2)
	for i := 0; i < nums1Length; i++ {
		nums1Map[nums1[i]] = 0
	}
	for i := 0; i < nums2Length; i++ {
		nums2Map[nums2[i]] = 0
	}
	for key := range nums1Map {
		if _, ok := nums2Map[key]; ok {
			result = append(result, key)
		}
	}
	return result
}

func removeOuterParentheses(S string) string {
	left, right, position, result := 0, 0, 0, ""
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			left += 1
		} else {
			right += 1
		}
		if left == right {
			result += S[position+1 : i]
			left, right = 0, 0
			position = i + 1
		}
	}
	return result
}

func main() {
	fmt.Println(diStringMatch("DDI"))
}
