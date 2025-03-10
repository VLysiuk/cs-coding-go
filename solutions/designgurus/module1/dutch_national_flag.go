package module1

/*
	Given an array containing 0s, 1s and 2s, sort the array in-place.
	You should treat numbers of the array as objects, hence, we can’t count 0s, 1s, and 2s to recreate the array.
	The flag of the Netherlands consists of three colors: red, white and blue;
	and since our input array also consists of three different numbers that is why it is called Dutch National Flag problem.
*/

// ==========Solution-1: Using one pointer and two cycles==========
// Time complexity: O(n)
// space: O(1)

func dutchNationalFlagSort(arr []int) []int {
	ptr := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			tmp := arr[ptr]
			arr[ptr] = arr[i]
			arr[i] = tmp
			ptr++
		}
	}

	for i := ptr; i < len(arr); i++ {
		if arr[i] == 1 {
			tmp := arr[ptr]
			arr[ptr] = arr[i]
			arr[i] = tmp
			ptr++
		}
	}

	return arr
}

// ==========Solution-2: Using two pointers==========
// Time complexity: O(n)
// space: O(1)

func dutchNationalFlagSort1(arr []int) []int {
	low, high := 0, len(arr)-1

	for i := 0; i <= high; {
		if arr[i] == 0 {
			arr[i], arr[low] = arr[low], arr[i]
			i++
			low++
		} else if arr[i] == 1 {
			i++
		} else {
			arr[i], arr[high] = arr[high], arr[i]
			high--
		}
	}

	return arr
}

// =================================================================================
