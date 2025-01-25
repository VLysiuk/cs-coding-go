package module1

/*
	Given a sorted array, create a new array containing squares of all the numbers of the input array in the sorted order.
*/
// ==========Solution-1: Using two pointers - moving to the edges==========
// Time complexity: O(n)
// space: O(n)

func squareSortedArray(arr []int) []int {
	n := len(arr)
	squares := make([]int, n)

	r, l := 0, 0

	for r < len(arr) && arr[r] < 0 {
		r++
	}

	l = r - 1

	for i := 0; i < len(arr); i++ {
		if l < 0 {
			squares[i] = arr[r] * arr[r]
			r++
		} else if r >= len(arr) {
			squares[i] = arr[l] * arr[l]
			l--
		} else if arr[l]*arr[l] < arr[r]*arr[r] {
			squares[i] = arr[l] * arr[l]
			l--
		} else {
			squares[i] = arr[r] * arr[r]
			r++
		}
	}

	return squares
}

// ==========Solution-2: Using two pointers - moving to the center==========
// Time complexity: O(n)
// space: O(n)

func squareSortedArray1(arr []int) []int {
	n := len(arr)
	squares := make([]int, n)

	l, r := 0, n-1
	sIndx := n - 1

	for l <= r {
		if arr[l]*arr[l] > arr[r]*arr[r] {
			squares[sIndx] = arr[l] * arr[l]
			l++
		} else {
			squares[sIndx] = arr[r] * arr[r]
			r--
		}
		sIndx--
	}

	return squares
}

// =================================================================================
