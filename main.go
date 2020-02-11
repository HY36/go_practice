package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	count := 0
	length := len(heights)
	for i := 1; i < length; i++ {
		if (heights[i] - heights[i-1]) > 1 {
			count++
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
	node.Val = node.Next.Val
	node.Next = node.Next.Next
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

func oddCells(n int, m int, indices [][]int) int {
	matrix, count := make([][]int, n), 0
	for i := 0; i < n; i++ {
		innerMatrix := make([]int, m)
		matrix[i] = innerMatrix
	}
	for _, out := range indices {
		for inIndex, in := range out {
			if inIndex == 0 {
				for i := 0; i < m; i++ {
					matrix[in][i]++
				}
			} else {
				for i := 0; i < n; i++ {
					matrix[i][in]++
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j]%2 == 1 {
				count++
			}
		}
	}
	return count
}

// NO.700 Search in a Binary Search Tree
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return root
	}
}

// NO.104 Maximum Depth of Binary Tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// NO.111 Minimum Depth of Binary Tree
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}

// NO.111 Minimum Depth of Binary Tree
// second solution
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDeepth := minDepth2(root.Left)
	rDeepth := minDepth2(root.Right)

	if lDeepth == 0 || rDeepth == 0 {
		return lDeepth + rDeepth + 1
	}
	return min(lDeepth, rDeepth) + 1
}

// min number
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// NO. 1313 Decompress Run-Length Encoded List
func decompressRLElist(nums []int) []int {
	nLength := len(nums)
	result := make([]int, 0)
	for i := 0; i < nLength-1; i += 2 {
		a, b := nums[i], nums[i+1]
		for j := 0; j < a; j++ {
			result = append(result, b)
		}
	}
	return result
}

// 1281. Subtract the Product and Sum of Digits of an Integer
func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for {
		if n == 0 {
			break
		}
		product *= n % 10
		sum += n % 10
		n /= 10
	}
	return product - sum
}

// 1295. Find Numbers with Even Number of Digits
func findNumbers(nums []int) int {
	var result, count int
	for _, n := range nums {
		count = 0
		for {
			if n == 0 {
				break
			}
			n /= 10
			count++
		}
		if count%2 == 0 {
			result++
		}
	}
	return result
}

// 1290. Convert Binary Number in a Linked List to Integer
func getDecimalValue(head *ListNode) int {
	container, result := make([]int, 0), 0
	for {
		container = append(container, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	length := len(container)
	for i := 0; i < length; i++ {
		result += container[i] << (length - i - 1)
	}
	return result
}

// 1323. Maximum 69 Number
func maximum69Number(num int) int {
	n := strconv.Itoa(num)
	n = strings.Replace(n,
		"6", "9", 1)
	result, _ := strconv.Atoi(n)
	return result
}

// 1304. Find N Unique Integers Sum up to Zero
func sumZero(n int) []int {
	if n == 1 {
		return []int{0}
	}
	container, sum := make([]int, n), 0
	for i := 0; i < n-1; i++ {
		container[i] = i + 1
		sum += i + 1
	}
	container[n-1] = -sum
	return container
}

// 938. Range Sum of BST
func rangeSumBST(root *TreeNode, L int, R int) int {
	return 0
}

// 1309.Decrypt String from Alphabet to Integer Mapping
func freqAlphabets(s string) string {
	length, i := len(s), 0
	result := make([]byte, 0)
	for {
		if i == length {
			break
		}
		if s[i] != '#' {
			result = append(result, s[i]+48)
			i++
		} else {
			i -= 2
			result = result[:len(result)-2]
			if s[i] == '1' {
				result = append(result, s[i]+s[i+1]+9)
			} else {
				result = append(result, s[i]+s[i+1]+18)
			}
			i += 3
		}

	}
	return string(result)
}

// 1207. Unique Number of Occurrences
func uniqueOccurrences(arr []int) bool {
	container := make(map[int]int)
	for i := range arr {
		if _, ok := container[arr[i]]; ok {
			container[arr[i]]++
		} else {
			container[arr[i]] = 1
		}
	}
	uniqueContainer := make(map[int]int)
	for _, v := range container {
		if _, ok := uniqueContainer[v]; ok {
			uniqueContainer[v]++
		} else {
			uniqueContainer[v] = 1
		}
	}
	if len(uniqueContainer) != len(container) {
		return false
	}
	return true
}

// todo: optimizer
// 1299. Replace Elements with Greatest Element on Right Side
func replaceElements(arr []int) []int {
	arrLength := len(arr)
	result := make([]int, arrLength)

	for i := 0; i < arrLength-1; i++ {
		tmp := 0
		for j := i + 1; j < arrLength; j++ {
			if arr[j] > tmp {
				tmp = arr[j]
			}
		}
		result[i] = tmp
	}
	result[arrLength-1] = -1
	return result
}

func rotate(nums []int, k int) {
	// length, tmp := len(nums), 0

}

// NO.26 Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			count++
			nums[count] = nums[i]
		}
	}
	return count + 1
}

// NO.121 Best Time to Buy and Sell Stock
func maxProfit1(prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 1 {
		return 0
	}
	profit, minPrice, currentProfit := 0, prices[0], 0
	for i := 1; i < pricesLen; i++ {
		currentProfit = prices[i] - minPrice
		if currentProfit < 0 {
			minPrice = prices[i]
		} else {
			if profit < currentProfit {
				profit = currentProfit
			}
		}
	}
	return profit
}

// NO.122 Best Time to Buy and Sell Stock II
func maxProfit2(prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 1 {
		return 0
	}
	profit, minPrice, lastProfit, currentProfit := 0, prices[0], 0, 0
	for index := 1; index < pricesLen; index++ {
		currentProfit = prices[index] - minPrice
		if currentProfit < lastProfit {
			if currentProfit >= 0 {
				profit += lastProfit
			}
			minPrice = prices[index]
		}
		lastProfit = currentProfit
	}
	if profit < (prices[pricesLen-1] - prices[0]) {
		profit = prices[pricesLen-1] - prices[0]
	}
	return profit
}

// 206. Reverse Linked List
func reverseList(head *ListNode) *ListNode {

	for {
		if head.Next == nil {
			break
		}

	}

	return head
}

func main() {
	input := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(input))
	input = []int{1, 2}
	fmt.Println(uniqueOccurrences(input))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
