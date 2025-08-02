package module2

/*
  Given the head of a Singly LinkedList that contains a cycle, write a function to find the starting node of the cycle.
*/

// ==========Solution-1: Using fast and slow pointers==========
// Time complexity: O(n^2)
// Space complexity: O(1)
func detectCycle(head *ListNode) *ListNode {
	cycleStart := head
	fast := head.Next
	slow := head.Next.Next

	for cycleStart != fast && cycleStart != slow {
		for fast != slow {
			slow = slow.Next
			fast = fast.Next.Next

			if cycleStart == fast || cycleStart == slow {
				return cycleStart
			}
		}
		cycleStart = cycleStart.Next
		slow = slow.Next
		fast = fast.Next.Next
	}

	return cycleStart
}

// ==========Solution-2: Using fast and slow pointers==========
// Time complexity: O(n)
// Space complexity: O(1)
func detectCycle1(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			break
		}
	}
	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
