package module2

import "strconv"

/*
  	Any number will be called a happy number if, after repeatedly replacing it with a number equal to the sum of the square of all of its digits, leads us to the number 1.
	All other (not-happy) numbers will never reach 1. Instead, they will be stuck in a cycle of numbers that does not include 1.
	Given a positive number n, return true if it is a happy number otherwise return false.
*/

// ==========Solution-1: Using map==========
// Time complexity: O(n)
// Space complexity: O(n)
func isHappy(num int) bool {
	nums := make(map[int]bool)
	result := num

	for result != 1 {
		result = squareNum(result)

		if nums[result] {
			return false
		}
		nums[result] = true
	}

	return true
}

func squareNum(num int) int {
	numStr := strconv.Itoa(num)
	var result int

	for _, c := range numStr {
		n := int(c - '0')
		result += n * n
	}

	return result
}

// =================================================================================
