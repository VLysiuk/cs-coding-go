package module1

import (
	"sort"
)

/*
	Given an array of unsorted numbers and a target number,
	find all unique quadruplets in it, whose sum is equal to the target number.
*/

// ==========Solution-1: Using two pointers==========
// Complexity: O(n^3)
// Space: O(n)
func quadrupleSumToTarget(arr []int, target int) [][]int {
	var quadruplets [][]int
	sort.Ints(arr)

	l, r := 0, 0

	for i := 0; i < len(arr)-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr)-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			l = j + 1
			r = len(arr) - 1

			for l < r {
				if arr[i]+arr[j]+arr[l]+arr[r] > target {
					r--
				} else if arr[i]+arr[j]+arr[l]+arr[r] < target {
					l++
				} else {
					quadruplets = append(quadruplets, []int{arr[i], arr[j], arr[l], arr[r]})
					l++
					r--
					for l < r && arr[l] == arr[l-1] {
						l++
					}
					for l < r && arr[r] == arr[r+1] {
						r--
					}
				}
			}
		}
	}
	return quadruplets
}

// =================================================================================
