package main

func twoSum(nums []int, target int) []int {
	// map[value]index from nums slice
	numMap := make(map[int]int, 0)

	for index, num := range nums {
		// 	expectedNum=target-num
		if val, present := numMap[target-num]; present {
			return []int{val, index}
		}

		// add record to map
		numMap[num] = index
	}
	return []int{}
}

func BestTimeToSellStock(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var maxProfit, minPrice int
	minPrice = prices[0]

	for _, price := range prices {
		// check whether there is any minimum
		if minPrice > price {
			minPrice = price
		} else if price-minPrice > maxProfit {
			// if not, then check for profit comparison
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// to track char & its frequency
	sMap := make(map[rune]int, 0)

	for _, ch := range s {
		sMap[ch]++
	}

	// tracking every char via sMap
	for _, ch := range t {
		sMap[ch]--
		if sMap[ch] < 0 {
			return false
		}
	}

	return true
}

func moveZeros(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}
}

// to check valid palindrome
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric from left
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}
		// Skip non-alphanumeric from right
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}
		// Compare lowercase versions
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// longestPrefix returns the longest common prefix string amongst an array of strings.
func longestPrefix(strs []string) string {
	/*
		*** Longest Prefix ***
		Logic:
		1. Take the first string as reference
		2. Compare each character of it with all other strings
		3. If mismatch found → return prefix till previous index
		4. If loop completes → first string itself is prefix

		input=>   strs: {"flower", "flow", "flight"}
		output=> "fl"
		comparison is done character-by-character from each string like:
		f == f == f ✔
		l == l == l ✔
		o == o != i ❌

		First loop-> to iterate through the each character of first string and comapre with other strings => row-wise traversing
		Second loop-> to iterate through each of the strings and do the comparison character-by-character => column-wise traversing
	*/
	if len(strs) == 0 {
		return ""
	}

	// traverse through each character of first string and compare with other strings
	for i := 0; i < len(strs[0]); i++ {
		currentChar := strs[0][i]

		// traverse through each string
		for j := 1; j < len(strs); j++ {

			// If index exceeds length OR mismatch found
			if (i > len(strs[j])) || (currentChar != strs[j][i]) {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[count-1] {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

// Merge sorted array
func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := 0

	for index, num := range nums2 {
		if num <= nums2[index] {
			nums1 = append(nums1[:pos], append([]int{num}, nums1[pos:]...)...)
			// pos++
		}
		pos++

	}
}
