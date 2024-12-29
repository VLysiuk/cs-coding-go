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
