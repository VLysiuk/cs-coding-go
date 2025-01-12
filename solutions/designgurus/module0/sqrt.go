package module0

/*
	Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
	The returned integer should be non-negative as well.
*/

// ==========Solution-1: Using binary search and recursion==========
// Time complexity: O(logn)
// space: O(1)

func mySqrt(x int) int {

	return findSqrt(1, x, x)
}

func findSqrt(low, high, x int) int {
	mid := (high + low) / 2

	if mid == low {
		return low
	}

	if mid*mid == x {
		return mid
	}

	if mid*mid < x {
		return findSqrt(mid, high, x)
	} else {
		return findSqrt(low, mid, x)
	}
}

// ==========Solution-2: Using binary search iterative==========
// Time complexity: O(logn)
// space: O(1)

func mySqrt1(x int) int {
	if x < 2 {
		return x
	}

	low, high := 1, x

	for low < high {
		mid := (high + low) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low - 1
}

// ==========Solution-3: Brute force==========
// Time complexity: O(n)
// space: O(1)

func mySqrt2(x int) int {
	for i := 1; i <= x; i++ {
		if i*i > x {
			return i - 1
		}
	}

	return x
}

// =================================================================================
