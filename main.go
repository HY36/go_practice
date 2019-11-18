package main

import (
	"fmt"
	"sort"
	"strings"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

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
			rightIndex--
		} else {
			result[index] = leftIndex
			leftIndex++
		}
		index++
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
			left++
		} else {
			right++
		}
		if left == right {
			result += S[position+1 : i]
			left, right = 0, 0
			position = i + 1
		}
	}
	return result
}
func longestPalindrome(s string) int {
	container := make(map[rune]int)
	even, odd, oddCount := 0, 0, 0
	for _, value := range s {
		container[value]++
	}
	if len(container) == 1 {
		return container[rune(s[0])]
	}
	for _, value := range container {
		if value%2 == 0 {
			even += value
		} else {
			odd += value
			oddCount++
		}
	}
	if odd != 0 {
		odd = odd - oddCount + 1
	}
	return even + odd
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
			count++
		} else {
			index++
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
						result++
						break
					}
					if board[i][m] == 'B' {
						break
					}
				}
				for m := j + 1; m < 8; m++ {
					if board[i][m] == 'p' {
						result++
						break
					}
					if board[i][m] == 'B' {
						break
					}
				}
				for n := i - 1; n >= 0; n-- {
					if board[n][j] == 'p' {
						result++
						break
					}
					if board[n][j] == 'B' {
						break
					}
				}
				for n := i + 1; n < 8; n++ {
					if board[n][j] == 'p' {
						result++
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

func commonChars(A []string) []string {
	result := make([]string, 0)
	return result

}

func isPalindrome(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}
	left, right, length := 0, len(s)-1, len(s)
	for left <= right {
		leftValue, rightValue := s[left], s[right]
		for !((leftValue <= 122 && leftValue >= 97) || (leftValue <= 90 && leftValue >= 65) || (leftValue <= 57 && leftValue >= 48)) {
			if left < length-1 {
				left++
				leftValue = s[left]
			} else {
				return true
			}

		}
		for !((rightValue <= 122 && rightValue >= 97) || (rightValue <= 90 && rightValue >= 65) || (rightValue <= 57 && rightValue >= 48)) {
			if right > 0 {
				right--
				rightValue = s[right]
			} else {
				return true
			}

		}
		if leftValue < 97 && !(leftValue < 58 && leftValue > 47) {
			leftValue += 32
		}
		if rightValue < 97 && !(rightValue < 58 && rightValue > 47) {
			rightValue += 32
		}
		if leftValue == rightValue {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func game(guess []int, answer []int) int {
	count := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			count++
		}
	}
	return count
}

func balancedStringSplit(s string) int {
	container := make([]string, 0, 0)
	count := 0
	container = append(container, string(s[0]))
	for i := 1; i < len(s); i++ {
		m := string(s[i])
		if len(container) == 0 {
			container = append(container, m)
			continue
		}
		currentVal := m + container[len(container)-1]
		if currentVal == "LR" || currentVal == "RL" {
			container = container[:len(container)-1]
			if len(container) == 0 {
				count++
			}
		} else {
			container = append(container, m)
		}
	}
	return count
}

func deleteNode(node *ListNode) {

}

func lengthOfLastWord(s string) int {
	lastVal, count := 0, 0
	for _, v := range s {
		if v == 32 {
			count = 0
		} else {
			count++
			lastVal = count
		}
	}
	return lastVal
}

func main() {
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord("Hello World!"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("b   a    "))
}
