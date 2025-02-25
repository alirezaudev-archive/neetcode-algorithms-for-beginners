package main

import "../../libs/Go/listnode"

func main() {
	reverseList(listnode.FromArray([]int{1, 2, 3, 4})).Print()
}

// https://leetcode.com/problems/design-browser-history/
func reverseList(head *listnode.ListNode) *listnode.ListNode {
	var prev *listnode.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
