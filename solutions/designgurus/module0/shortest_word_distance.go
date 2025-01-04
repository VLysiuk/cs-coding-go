package module0

/*
	Given an array of strings words and two different strings that already exist in the array word1 and word2,
	return the shortest distance between these two words in the list.
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n)
// space: O(1)

func shortestWordDistance(words []string, word1 string, word2 string) int {
	if len(words) < 2 {
		return 0
	}

	minDistance := len(words)
	w1_i, w2_i := -1, -1

	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			w1_i = i
		}

		if words[i] == word2 {
			w2_i = i
		}

		if w1_i != -1 && w2_i != -1 {
			distance := Abs(w2_i - w1_i)

			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}

	return x
}

// =================================================================================
