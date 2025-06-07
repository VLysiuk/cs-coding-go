package module1

/*
	Given an array, find the length of the smallest subarray in it which when sorted will sort the whole array.
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n)
// Space: O(1)
func minimumWindowSort(arr []int) int {
	start, end := 0, 0
	min, max := arr[0], arr[0]
	//find min and max
	for _, v := range arr {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	//find start
	if arr[0] > min {
		start = 0
	} else {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] >= arr[i+1] {
				start = i
				break
			}
		}
	}

	//find end
	if arr[len(arr)-1] < max {
		end = len(arr) - 1
	} else {
		for i := len(arr) - 1; i > 0; i-- {
			if arr[i] < arr[i-1] {
				end = i
				break
			}
		}
	}

	if start < end {
		return end - start + 1
	}

	return 0
}

// =================================================================================
