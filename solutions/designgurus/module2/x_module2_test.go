package module2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListCycle(t *testing.T) {
	tests := []struct {
		head     *ListNode
		expected bool
	}{
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}, false},
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}, false},
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}, false},
		{&ListNode{Val: 1, Next: nil}, false},
	}

	for _, test := range tests {
		result := hasCycle(test.head)
		assert.Equal(t, test.expected, result)
	}

	testsWithCycle := []struct {
		head     *ListNode
		expected bool
	}{
		{
			func() *ListNode {
				head := &ListNode{Val: 1}
				second := &ListNode{Val: 2}
				third := &ListNode{Val: 3}
				head.Next = second
				second.Next = third
				third.Next = second // creates a cycle
				return head
			}(),
			true,
		},
		{
			func() *ListNode {
				head := &ListNode{Val: 1}
				second := &ListNode{Val: 2}
				head.Next = second
				second.Next = head // creates a cycle
				return head
			}(),
			true,
		},
	}

	for _, test := range testsWithCycle {
		result := hasCycle(test.head)
		assert.Equal(t, test.expected, result)
	}
}
