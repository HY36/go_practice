package main

import (
	"fmt"
	"sort"
	"strings"
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

func defangIPaddr(address string) string {
	length := len(address)
	result := make([]string, length+8)
	for i := 0; i < length; i++ {
		if address[i] != '.' {
			result = append(result, string(address[i]))
		} else {
			result = append(result, "[.]")
		}
	}
	return strings.Join(result, "")

}

func heightChecker(heights []int) int {
	index := 0
	count := 0
	length := len(heights)
	for i := 1; i < length; i++ {
		if !(heights[i] >= heights[index]) {
			count += 1
		} else {
			index += 1
		}
	}
	return count
}

func findOcurrences(text string, first string, second string) []string {
	textList := strings.Split(text, " ")
	length := len(textList)
	result := make([]string, 0)
	for i := 0; i < length-2; i++ {
		if textList[i] == first && textList[i+1] == second {
			result = append(result, textList[i+2])
		}
	}
	return result
}

func numRookCaptures(board [][]byte) int {
	result := 0
	for i, iv := range board {
		for j, jv := range iv {
			if jv == 'R' {
				for m := j - 1; m >= 0; m-- {
					if board[i][m] == 'p' {
						result += 1
						break
					}
					if board[i][m] == 'B' {
						break
					}
				}
				for m := j + 1; m < 8; m++ {
					if board[i][m] == 'p' {
						result += 1
						break
					}
					if board[i][m] == 'B' {
						break
					}
				}
				for n := i - 1; n >= 0; n-- {
					if board[n][j] == 'p' {
						result += 1
						break
					}
					if board[n][j] == 'B' {
						break
					}
				}
				for n := i + 1; n < 8; n++ {
					if board[n][j] == 'p' {
						result += 1
						break
					}
					if board[n][j] == 'B' {
						break
					}
				}
			}
		}
	}
	return result
}

func main() {
	//fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	input := [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', 'p', 'p', '.', '.', '.', '.'}, {'.', 'p', 'p', 'R', '.', 'p', '.', 'p'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', 'p', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(numRookCaptures(input))
}
