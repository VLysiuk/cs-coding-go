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

// =================================================================================
