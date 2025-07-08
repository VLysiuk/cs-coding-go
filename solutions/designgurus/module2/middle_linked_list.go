package module2

/*
	Given the head of a Singly LinkedList, write a method to return the middle node of the LinkedList.
	If the total number of nodes in the LinkedList is even, return the second middle node.
*/

// ==========Solution-1: Using fast and slow pointers==========
// Time complexity: O(n)
// Space complexity: O(1)
func findMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow_ptr := head
	fast_ptr := head.Next

	for fast_ptr != nil {
		slow_ptr = slow_ptr.Next
		fast_ptr = fast_ptr.Next
		if fast_ptr != nil {
			fast_ptr = fast_ptr.Next
		}
	}
	return slow_ptr
}
