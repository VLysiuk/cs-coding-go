package module0

/*
	Given an array of integers nums, return the number of good pairs.
	A pair (i, j) is called good if nums[i] == nums[j] and i < j.
*/

// ==========Solution-1: Using a map and ariphmetic progression formula==========
// Time complexity: O(n)
// space: O(n)

func numGoodPairs(nums []int) int {
	pairCount := 0

	pairMap := make(map[int]int)

	for _, v := range nums {
		pairMap[v] += 1
	}

	for _, v := range pairMap {
		pairCount += v * (v - 1) / 2
	}

	return pairCount
}

// ==========Solution-2: Using a map to store the frequency of each element==========
// Time complexity: O(n)
// space: O(n)

func numGoodPairs1(nums []int) int {
	pairCount := 0

	pairMap := make(map[int]int)

	for _, v := range nums {
		pairMap[v]++

		pairCount += pairMap[v] - 1
	}

	return pairCount
}

// ==========Solution-3: Brute force===
// Time complexity: O(n^2)
// space: O(1)

func numGoodPairs2(nums []int) int {
	pairCount := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				pairCount++
			}
		}
	}

	return pairCount
}

// =================================================================================
