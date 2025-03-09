package module1

import (
	"sort"
)

/*
	Given an array arr of unsorted numbers and a target sum,
	count all triplets in it such that arr[i] + arr[j] + arr[k] < target where i, j, and k are three different indices.
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n^2)
// space: O(n)

func tripletsWithSmallerSum(arr []int, target int) int {
	count := 0
	sort.Ints(arr)
	l, r := 0, 0

	for i := 0; i < len(arr)-2; i++ {
		l = i + 1
		r = len(arr) - 1
		for l < r {
			if arr[i]+arr[l]+arr[r] < target {
				count += r - l
				l++
			} else {
				r--
			}
		}
	}

	return count
}

// =================================================================================
