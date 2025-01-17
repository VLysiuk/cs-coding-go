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
