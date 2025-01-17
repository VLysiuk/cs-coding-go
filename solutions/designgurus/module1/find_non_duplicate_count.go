package module1

/*
	Given an array of sorted numbers, move all non-duplicate number instances at the beginning of the array in-place.
	The non-duplicate numbers should be sorted and you should not use any extra space so that the solution has constant space complexity i.e., .
	Move all the unique number instances at the beginning of the array and after moving return the length of the subarray that has no duplicate in it.
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n)
// space: O(1)

func findNonDuplicateCount(arr []int) int {
	current := 0
	next := 1

	for next < len(arr) {
		for next < len(arr) && arr[current] >= arr[next] {
			next++
		}

		if next == len(arr) {
			return current + 1
		}

		tmp := arr[current+1]
		arr[current+1] = arr[next]
		arr[next] = tmp

		current++
		next++
	}

	return current + 1
}

// =================================================================================
