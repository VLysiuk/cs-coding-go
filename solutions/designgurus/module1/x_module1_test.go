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
