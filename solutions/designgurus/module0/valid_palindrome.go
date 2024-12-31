package module0

import "unicode"

/*
	A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
	Given a string s, return true if it is a palindrome, or false otherwise.
*/

// ==========Solution-1: Going trough array from both sides==========
// Time complexity: O(n)
// space: O(n)

func isPalindrome(s string) bool {
	cleanStr := make([]rune, 0, len(s))

	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			cleanStr = append(cleanStr, unicode.ToLower(c))
		}
	}

	strLength := len(cleanStr)
	for i := 0; i < strLength; i++ {
		if cleanStr[i] != cleanStr[strLength-i-1] {
			return false
		}
	}

	return true
}

// ==========Solution-2: Using two pointers===
// Time complexity: O(n)
// space: O(1)

func isPalindrome1(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isLetterOrDigit(rune(s[left])) {
			left++
		}

		for left < right && !isLetterOrDigit(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}

func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// =================================================================================
