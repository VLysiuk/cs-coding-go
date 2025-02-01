package module1

import (
	"math"
	"sort"
)

/*
	Given an array of unsorted numbers and a target number,
	find a triplet in the array whose sum is as close to the target number as possible,
	return the sum of the triplet. If there are more than one such triplet, return the sum of the triplet with the smallest sum.
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n^2)
// space: O(n)

func tripletsSumToTarget(arr []int, targetSum int) int {
	sort.Ints(arr)
	closestSum := 0
	closestDistance := math.MaxInt32
	l, r := 0, 0

	for i := 0; i < len(arr)-2; i++ {
		l = i + 1
		r = len(arr) - 1

		for l < r {
			sum := arr[i] + arr[l] + arr[r]
			distance := Abs(targetSum - sum)
			if distance < closestDistance {
				closestDistance = distance
				closestSum = sum
			} else if distance == closestDistance && sum < closestSum {
				closestDistance = distance
				closestSum = sum
			}

			if sum == targetSum {
				return sum
			}

			if sum < targetSum {
				l++
			} else {
				r--
			}
		}

	}

	return closestSum
}

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// =================================================================================
