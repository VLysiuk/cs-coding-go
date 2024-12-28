package main

import "fmt"

/*
	Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/

// complexity: O(n)
// space: O(n)
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = 1
		}
	}

	return false
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{1, 2, 3, 1}
	nums3 := []int{-1, 2, 3, -1}
	nums4 := []int{-1, -2, 0, 0, 1, 5, 6}
	nums5 := []int{-1, -2, 0, 1, 5}

	fmt.Println(containsDuplicate(nums1)) // false
	fmt.Println(containsDuplicate(nums2)) // true
	fmt.Println(containsDuplicate(nums3)) // true
	fmt.Println(containsDuplicate(nums4)) // true
	fmt.Println(containsDuplicate(nums5)) // false
}
