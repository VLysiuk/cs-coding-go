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

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"hello", "world", false},
		{"listen", "silent", true},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"rat", "tar", true},
		{"aacc", "caac", true},
		{"rat", "cat", false},
	}

	for _, test := range tests {
		result := isAnagram(test.s, test.t)
		result1 := isAnagram1(test.s, test.t)
		result2 := isAnagram2(test.s, test.t)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expected, result1)
		assert.Equal(t, test.expected, result2)
	}
}

func TestShortestWordDistance(t *testing.T) {
	tests := []struct {
		words       []string
		word1       string
		word2       string
		minDistance int
	}{
		{[]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice", 3},
		{[]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding", 1},
		{[]string{"practice", "makes", "perfect", "coding", "makes"}, "practice", "makes", 1},
		{[]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "makes", 1},
		{[]string{}, "coding", "perfect", 0},
		{[]string{"a", "b", "c", "d", "e"}, "a", "e", 4},
		{[]string{"a", "c", "d", "b", "a", "c", "c", "a", "d", "b"}, "a", "b", 1},
		{[]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}, "fox", "dog", 5},
	}

	for _, test := range tests {
		result := shortestWordDistance(test.words, test.word1, test.word2)
		assert.Equal(t, test.minDistance, result)
	}
}

func TestNumGoodPairs(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
		{[]int{1, 2, 3, 4, 5, 6}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5}, 5},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 45},
	}

	for _, test := range tests {
		result := numGoodPairs(test.nums)
		result1 := numGoodPairs1(test.nums)
		result2 := numGoodPairs2(test.nums)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expected, result1)
		assert.Equal(t, test.expected, result2)
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
		{16, 4},
		{25, 5},
		{36, 6},
		{49, 7},
		{64, 8},
		{81, 9},
		{100, 10},
		{121, 11},
		{144, 12},
		{169, 13},
		{196, 14},
		{225, 15},
		{256, 16},
		{289, 17},
		{484, 22},
		{529, 23},
		{576, 24},
		{625, 25},
		{676, 26},
		{729, 27},
		{784, 28},
		{841, 29},
		{900, 30},
		{961, 31},
		{1024, 32},
		{1089, 33},
		{1156, 34},
		{1225, 35},
		{1296, 36},
		{1369, 37},
		{1444, 38},
		{1521, 39},
		{1600, 40},
		{2500, 50},
		{3025, 55},
		{3136, 56},
		{3249, 57},
	}

	for _, test := range tests {
		result := mySqrt(test.x)
		result1 := mySqrt1(test.x)
		result2 := mySqrt2(test.x)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expected, result1)
		assert.Equal(t, test.expected, result2)
	}
}
