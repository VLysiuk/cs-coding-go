package module0

import "unicode"

/*
	A pangram is a sentence where every letter of the English alphabet appears at least once.
	Given a string sentence containing English letters (lower or upper-case), return true if sentence is a pangram, or false otherwise.
	Note: The given sentence might contain other characters like digits or spaces, your solution should handle these too.
*/

// ==========Solution-1: Using a map to store a count of each letter==========
// Time complexity: O(n)
// space: O(1) - since the alphabet size is fixed

func isPangram(sentence string) bool {
	alphabet := map[rune]int{
		'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0,
		'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0,
		'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0,
	}

	for _, r := range sentence {
		if _, exists := alphabet[unicode.ToLower(r)]; exists {
			alphabet[unicode.ToLower(r)] = 1
		}
	}

	var totalSum = 0
	for _, v := range alphabet {
		totalSum += v
	}

	return totalSum >= 26
}

// ==========Solution-2: Using an empty map(HashSet) and isLetter method==========
// Time complexity: O(n)
// space: O(1) - since the alphabet size is fixed

func isPangram1(sentence string) bool {
	alphabet := make(map[rune]struct{})

	for _, r := range sentence {
		if unicode.IsLetter(r) {
			alphabet[unicode.ToLower(r)] = struct{}{}
		}
	}

	return len(alphabet) == 26
}

// =================================================================================
