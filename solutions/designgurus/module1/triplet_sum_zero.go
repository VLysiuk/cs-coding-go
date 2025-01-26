package module1

import "sort"

/*
	Given an array of unsorted numbers, find all unique triplets in it that add up to zero.
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n^2)
// space: O(n)

func searchTriplets(arr []int) [][]int {
	triplets := make([][]int, 0)
	sort.Ints(arr)

	l, r := 0, 0
	for i := 0; i < len(arr)-2; i++ {
		l = i + 1
		r = len(arr) - 1

		if i != 0 && arr[i] == arr[i-1] {
			continue // skip duplicates
		}
		for l < r {
			if arr[l]+arr[r]+arr[i] == 0 {
				triplets = append(triplets, []int{arr[i], arr[l], arr[r]})
				r--
				l++
				for l < r && arr[l] == arr[l-1] {
					l++ // skip same element to avoid duplicate triplets
				}
				for l < r && arr[r] == arr[r+1] {
					r-- // skip same element to avoid duplicate triplets
				}
			} else if arr[l]+arr[r]+arr[i] > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return triplets
}

// =================================================================================
