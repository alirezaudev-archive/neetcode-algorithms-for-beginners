package main

import "fmt"

// https://leetcode.com/problems/merge-k-sorted-lists/
func main() {
	lists := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	var linkedLists []*ListNode
	for _, list := range lists {
		linkedLists = append(linkedLists, FromArray(list))
	}

	mergeKLists(linkedLists).Print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	head := lists[0]
	for _, list := range lists[1:] {
		head = merge(head, list)
	}

	return head
}

func merge(left, right *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	leftCurr := left
	rightCurr := right
	for leftCurr != nil && rightCurr != nil {
		if leftCurr.Val < rightCurr.Val {
			curr.Next = leftCurr
			leftCurr = leftCurr.Next
		} else {
			curr.Next = rightCurr
			rightCurr = rightCurr.Next
		}
		curr = curr.Next
	}

	if leftCurr != nil {
		curr.Next = leftCurr
	}

	if rightCurr != nil {
		curr.Next = rightCurr
	}
	return head.Next
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
