package module0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{-1, 2, 3, -1}, true},
		{[]int{-1, -2, 0, 0, 1, 5, 6}, true},
		{[]int{-1, -2, 0, 1, 5}, false},
	}

	for _, test := range tests {
		result := containsDuplicate(test.nums)
		assert.Equal(t, test.expected, result)
	}
}

func TestIsPangram(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"The quick brown fox jumps over the lazy dog!1", true},
		{"TheQuickBrownFoxJumpsOverTheLazyDog", true},
		{"This is not a pangram", false},
		{"Hello World!", false},
		{"12 Pack my box with five dozen liquor jugs:123", true},
		{"", false},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"abcdefghijklmnopqrstuvwxyz", true},
	}

	for _, test := range tests {
		result := isPangram(test.input)
		result1 := isPangram1(test.input)
		assert.Equal(t, test.expected, result, "For input '%s'", test.input)
		assert.Equal(t, test.expected, result1, "For input '%s'", test.input)
	}
}

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
		{"example", "example"},
		{"", ""},
		{"aeiou", "uoiea"},
		{"racecar", "racecar"},
		{"AEIOU", "UOIEA"},
	}

	for _, test := range tests {
		result := reverseVowels(test.input)
		assert.Equal(t, test.expected, result, "For input '%s'", test.input)
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"madam", true},
		{"step on no pets", true},
		{"hello", false},
		{"", true},
		{"A man, a plan, a canal, Panama", true},
		{"No lemon, no melon", true},
		{"Was it a car or a cat I saw?", true},
		{"12321", true},
		{"12345", false},
		{"A man, a plan, a canal, Panama!", true},
		{"Was it a car or a cat I saw?", true},
		{"Hello World!", false},
		{"!!??-@$%&", true},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		result1 := isPalindrome1(test.input)
		assert.Equal(t, test.expected, result, "For input '%s'", test.input)
		assert.Equal(t, test.expected, result1, "For input '%s'", test.input)
	}
}
