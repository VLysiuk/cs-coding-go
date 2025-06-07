package module1

/*
	Given two strings containing backspaces (identified by the character ‘#’), check if the two strings are equal.
*/
// ==========Solution-1: Using stack==========
// Time complexity: O(n)
// Space: O(n)
func compareStringsWithBackspaces(str1, str2 string) bool {
	stack1 := cleanUpString(str1)
	stack2 := cleanUpString(str2)

	if len(stack1) != len(stack2) {
		return false
	}

	for i := 0; i < len(stack1); i++ {
		if stack1[i] != stack2[i] {
			return false
		}
	}

	return true
}

func cleanUpString(str string) []rune {
	stack := []rune{}

	for _, c := range str {
		if c != '#' {
			stack = append(stack, c)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return stack
}

// ==========Solution-2: Using two pointers==========
// Time complexity: O(n)
// Space: O(1)
func compareStringsWithBackspaces1(str1, str2 string) bool {
	index1 := len(str1) - 1
	index2 := len(str2) - 1

	for index1 >= 0 || index2 >= 0 {
		index1 = findNextValidCharIndex(str1, index1)
		index2 = findNextValidCharIndex(str2, index2)

		if index1 < 0 && index2 < 0 {
			return true // both strings are equal
		}

		if index1 < 0 || index2 < 0 {
			return false // one string is shorter than the other
		}

		if str1[index1] != str2[index2] {
			return false // characters do not match
		}

		index1--
		index2--
	}
	return true // all characters matched
}

func findNextValidCharIndex(str string, index int) int {
	backspaceCount := 0

	for index >= 0 {
		if str[index] == '#' {
			backspaceCount++
		} else if backspaceCount > 0 {
			backspaceCount--
		} else {
			return index
		}
		index--
	}
	return index
}

// =================================================================================
