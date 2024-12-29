package module0

/*
	Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/

// ==========Solution-1: Using a map to store the frequency of each element==========
// Time complexity: O(n)
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

// =================================================================================
