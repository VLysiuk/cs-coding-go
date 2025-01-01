package module0

import "sort"

/*
	Given two strings s and t, return true if t is an anagram of s, and false otherwise.
	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

// ==========Solution-1: Using a map to store the existance of each character==========
// Time complexity: O(n)
// space: O(n)

func isAnagram(s string, t string) bool {
	stringMap := make(map[rune]struct{})

	for _, r := range s {
		stringMap[r] = struct{}{}
	}

	for _, c := range t {
		if _, exists := stringMap[c]; !exists {
			return false
		}
	}

	return true
}

// ==========Solution-2: Brute force===
// Time complexity: O(n^2)
// space: O(1)

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	for _, c := range t {
		exists := false

		for _, r := range s {
			if c == r {
				exists = true
				break
			}
		}

		if !exists {
			return false
		}
	}

	return true
}

// ==========Solution-3: Sorting the strings and comparing them===
// Time complexity: O(nlogn)
// space: O(n)

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charsS := []rune(s)
	charsT := []rune(t)

	sort.Slice(charsS, func(i, j int) bool {
		return charsS[i] < charsS[j]
	})

	sort.Slice(charsT, func(i, j int) bool {
		return charsT[i] < charsT[j]
	})

	for i := 0; i < len(charsS); i++ {
		if charsS[i] != charsT[i] {
			return false
		}
	}

	return true
}

// =================================================================================
