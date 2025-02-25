package listnode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromArray(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	curr := head

	for _, v := range arr[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}

	return head
}

func (head *ListNode) Print() {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

func (head *ListNode) Reverse() *ListNode {
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
