package main

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newHead
}
