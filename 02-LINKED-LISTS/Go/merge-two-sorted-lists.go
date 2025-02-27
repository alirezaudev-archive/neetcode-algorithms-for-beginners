package main

import "../../libs/Go/linkedlist"

func main() {
	list1 := linkedlist.FromArray([]int{1, 2, 4})
	list2 := linkedlist.FromArray([]int{1, 3, 4})

	mergeTwoLists2(list1.Head, list2.Head).Print()
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *linkedlist.Node, list2 *linkedlist.Node) *linkedlist.Node {
	result := linkedlist.Node{}

	curr1 := list1
	curr2 := list2
	curr := &result
	for curr1 != nil && curr2 != nil {
		if curr2.Val < curr1.Val {
			curr.Next = curr2
			curr2 = curr2.Next
		} else {
			curr.Next = curr1
			curr1 = curr1.Next
		}

		curr = curr.Next
	}

	if curr1 != nil {
		curr.Next = curr1
	} else {
		curr.Next = curr2
	}

	return result.Next
}

// Solution 2
func mergeTwoLists2(list1 *linkedlist.Node, list2 *linkedlist.Node) *linkedlist.Node {
	if list2 != nil && (list1 == nil || list2.Val < list1.Val) {
		list1, list2 = list2, list1
	}

	curr1 := list1
	curr2 := list2

	for curr1 != nil && curr2 != nil {
		for curr1.Next != nil && curr1.Next.Val < curr2.Val {
			curr1 = curr1.Next
		}

		tmp := curr2.Next
		curr2.Next = curr1.Next
		curr1.Next = curr2

		curr2 = tmp
		curr1 = curr1.Next
	}

	return list1
}
