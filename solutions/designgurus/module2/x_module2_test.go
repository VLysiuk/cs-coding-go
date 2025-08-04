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

func TestFindMiddle(t *testing.T) {
	tests := []struct {
		head     *ListNode
		expected int
	}{
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}, 2},
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}, 3},
		{&ListNode{Val: 1}, 1},
	}

	for _, test := range tests {
		middle := findMiddle(test.head)
		assert.Equal(t, test.expected, middle.Val)
	}
}
func TestDetectCycleStart(t *testing.T) {
	tests := []struct {
		head     *ListNode
		expected int
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
			2, // starting node of the cycle
		},
		{
			func() *ListNode {
				head := &ListNode{Val: 1}
				second := &ListNode{Val: 2}
				head.Next = second
				second.Next = head // creates a cycle
				return head
			}(),
			1, // starting node of the cycle
		},
	}

	for _, test := range tests {
		cycleStart := detectCycle(test.head)
		cycleStart1 := detectCycle1(test.head)
		assert.Equal(t, test.expected, cycleStart.Val)
		assert.Equal(t, test.expected, cycleStart1.Val)
	}
}

func TestIsHappyNumber(t *testing.T) {
	tests := []struct {
		num      int
		expected bool
	}{
		{19, true},
		{2, false},
		{7, true},
		{1, true},
		{0, false},
		{23, true},
		{12, false},
	}

	for _, test := range tests {
		result := isHappy(test.num)
		assert.Equal(t, test.expected, result)
	}
}
