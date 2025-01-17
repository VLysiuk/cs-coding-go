package module1

/*
	Given an unsorted array of numbers and a target ‘key’, remove all instances of ‘key’ in-place and return the new length of the array.
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n)
// space: O(1)

func removeKey(arr []int, key int) int {
	insrtIdx := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != key {
			arr[insrtIdx] = arr[i]
			insrtIdx++
		}
	}

	return insrtIdx
}

// =================================================================================
