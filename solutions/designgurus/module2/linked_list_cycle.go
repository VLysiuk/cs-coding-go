package module2

/*
  Given the head of a Singly LinkedList, write a function to determine if the LinkedList has a cycle in it or not.
*/

// ListNode represents a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ==========Solution-2: Using fast and slow pointers==========
// Time complexity: O(n)
// Space complexity: O(1)
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// =================================================================================
