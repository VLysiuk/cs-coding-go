package module0

/*
	Given a string s, reverse only all the vowels in the string and return it.
	The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
*/

// ==========Solution-1: Using two pointers==========
// Time complexity: O(n)
// space: O(n)

func reverseVowels(str string) string {
	runeStr := []rune(str)
	left := 0
	right := len(str) - 1

	for left < right {
		for !isVowel(runeStr[left]) && left < right {
			left++
		}

		for !isVowel(runeStr[right]) && right > left {
			right--
		}

		if isVowel(runeStr[left]) && isVowel(runeStr[right]) {
			tmp := runeStr[left]
			runeStr[left] = runeStr[right]
			runeStr[right] = tmp
			left++
			right--
		}
	}

	return string(runeStr)
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
