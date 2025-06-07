package module1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairWithTargetSum(t *testing.T) {
	tests := []struct {
		arr       []int
		targetSum int
		expected  []int
	}{
		{[]int{1, 2, 3, 4, 6}, 6, []int{1, 3}},
		{[]int{2, 5, 9, 11}, 11, []int{0, 2}},
		{[]int{2, 3, 4}, 6, []int{0, 2}},
		{[]int{2, 3, 4, 7}, 6, []int{0, 2}},
		{[]int{2, 3, 4, 7}, 10, []int{1, 3}},
		{[]int{1, 2, 3, 4, 5, 9}, 18, []int{-1, -1}},
	}

	for _, test := range tests {
		result := pairWithTargetSum(test.arr, test.targetSum)
		result1 := pairWithTargetSum1(test.arr, test.targetSum)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expected, result1)
	}
}

func TestNonDuplicateCount(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 3, 3, 3, 6, 9, 9}, 4},
		{[]int{2, 2, 2, 11}, 2},
		{[]int{1, 2, 2, 2}, 2},
		{[]int{1, 2, 3, 3, 3, 4, 4, 5, 5}, 5},
		{[]int{2}, 1},
		{[]int{2, 2}, 1},
		{[]int{2, 3}, 2},
		{[]int{2, 3, 4}, 3},
	}

	for _, test := range tests {
		result := findNonDuplicateCount(test.arr)
		assert.Equal(t, test.expected, result)
	}
}

func TestRemoveKey(t *testing.T) {
	tests := []struct {
		arr      []int
		key      int
		expected int
	}{
		{[]int{3, 2, 3, 6, 3, 10, 9, 3}, 3, 4},
		{[]int{2, 11, 2, 2, 1}, 2, 2},
	}

	for _, test := range tests {
		result := removeKey(test.arr, test.key)
		assert.Equal(t, test.expected, result)
	}
}

func TestSquareSortedArray(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{-2, -1, 0, 2, 3}, []int{0, 1, 4, 4, 9}},
		{[]int{-3, -1, 0, 1, 2}, []int{0, 1, 1, 4, 9}},
		{[]int{-3, -2, -1}, []int{1, 4, 9}},
		{[]int{1, 2, 3}, []int{1, 4, 9}},
		{[]int{0, 1, 2, 3}, []int{0, 1, 4, 9}},
		{[]int{-3, -2, -1, 0}, []int{0, 1, 4, 9}},
	}

	for _, test := range tests {
		result := squareSortedArray(test.arr)
		result1 := squareSortedArray1(test.arr)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expected, result1)
	}
}

func TestTripletSum(t *testing.T) {
	tests := []struct {
		arr      []int
		expected [][]int
	}{
		{[]int{-3, 0, 1, 2, -1, 1, -2}, [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}}},
		{[]int{-5, 2, -1, -2, 3}, [][]int{{-5, 2, 3}, {-2, -1, 3}}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
	}

	for _, test := range tests {
		result := searchTriplets(test.arr)
		assert.ElementsMatch(t, test.expected, result)
	}
}

func TestTripletSumCloseToTarget(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{-2, 0, 1, 2}, 2, 1},
		{[]int{-3, -1, 1, 2}, 1, 0},
		{[]int{1, 0, 1, 1}, 100, 3},
		{[]int{1, 2, 3, 4, 5}, 100, 12},
		{[]int{1, 2, 3, 4, 5}, 10, 10},
		{[]int{-1, 0, 2, 3}, 3, 2},
		{[]int{39, -55, 11, 69, 4, -9, 6, 23}, -72, -60},
	}

	for _, test := range tests {
		result := tripletsSumToTarget(test.arr, test.target)
		assert.Equal(t, test.expected, result)
	}
}

func TestTripletWithSmallerSum(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{-1, 0, 2, 3}, 3, 2},
		{[]int{-1, 4, 2, 1, 3}, 5, 4},
		{[]int{7, 4, 1, 2, 3}, 8, 2},
		{[]int{1, 2, 3, 4, 5}, 10, 6},
		{[]int{1, 2, 3, 4, 5}, 100, 10},
		{[]int{2, 3, 4, 5}, 1, 0},
	}

	for _, test := range tests {
		result := tripletsWithSmallerSum(test.arr, test.target)
		assert.Equal(t, test.expected, result)
	}
}

func TestDutchNationalFlagSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{1, 0, 2, 1, 0}, []int{0, 0, 1, 1, 2}},
		{[]int{2, 2, 0, 1, 2, 0}, []int{0, 0, 1, 2, 2, 2}},
		{[]int{2, 2, 0, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
		{[]int{1, 1, 1, 0, 1}, []int{0, 1, 1, 1, 1}},
	}

	for _, test := range tests {
		result := dutchNationalFlagSort(test.arr)
		result1 := dutchNationalFlagSort1(test.arr)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expected, result1)
	}
}

func TestQuadrupleSumToTarget(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected [][]int
	}{
		{[]int{4, 1, 2, -1, 1, -3}, 1, [][]int{{-3, -1, 1, 4}, {-3, 1, 1, 2}}},
		{[]int{2, 0, -1, 1, -2, 2}, 2, [][]int{{-2, 0, 2, 2}, {-1, 0, 1, 2}}},
		{[]int{2, 0, -1, 1, -2, 2}, 0, [][]int{{-2, -1, 1, 2}}},
		{[]int{1, 2, 3, 4, 5}, 10, [][]int{{1, 2, 3, 4}}},
		{[]int{0, 0, 0, 0, 0, 0}, 0, [][]int{{0, 0, 0, 0}}},
	}

	for _, test := range tests {
		result := quadrupleSumToTarget(test.arr, test.target)
		assert.ElementsMatch(t, test.expected, result)
	}
}

func TestStringsWithBackspaces(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"xy#z", "xzz#", true},
		{"xy#z", "xyz#z#", false},
		{"xp#", "xyz##", true},
		{"xywrrmp", "xywrrmu#", false},
		{"a#c", "b", false},
		{"#", "#", true},
		{"a##c", "#a#c", true},
		{"ac###", "b#c#", true},
	}

	for _, test := range tests {
		result := compareStringsWithBackspaces(test.str1, test.str2)
		result1 := compareStringsWithBackspaces1(test.str1, test.str2)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expected, result1)
	}
}

func TestMinimumWindowSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 3, 2, 0, -1, 7, 10}, 5},
		{[]int{1, 2, 3}, 0},
		{[]int{3, 2, 1}, 3},
		{[]int{1, 3, 3, 2}, 3},
	}

	for _, test := range tests {
		result := minimumWindowSort(test.arr)
		assert.Equal(t, test.expected, result)
	}
}
