package module1

/*
	Given an array of numbers sorted in ascending order and a target sum, find a pair in the array whose sum is equal to the given target.
	Write a function to return the indices of the two numbers (i.e. the pair) such that they add up to the given target. If no such pair exists return [-1, -1].
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n)
// space: O(1)

func pairWithTargetSum(arr []int, targetSum int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		if arr[left]+arr[right] == targetSum {
			return []int{left, right}
		}

		if arr[left]+arr[right] > targetSum {
			right--
		} else {
			left++
		}
	}

	return []int{-1, -1} // pair not found
}

// ==========Solution-2: Using a map==========
// Time complexity: O(n)
// space: O(n)

func pairWithTargetSum1(arr []int, targetSum int) []int {
	m := make(map[int]int)

	for i, v := range arr {
		if v1, ok := m[targetSum-v]; ok {
			return []int{v1, i}
		} else {
			m[v] = i
		}
	}

	return []int{-1, -1}
}

// =================================================================================
