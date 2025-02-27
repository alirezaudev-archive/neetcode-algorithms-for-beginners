package main

type ListNode struct {
	Next *ListNode
	Prev *ListNode
}

// https://leetcode.com/problems/design-browser-history/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
