package main

import "../../libs/Go/listnode"

func main() {
	list1 := listnode.FromArray([]int{1, 2, 4})
	list2 := listnode.FromArray([]int{1, 3, 4})

	mergeTwoLists2(list1, list2).Print()
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *listnode.ListNode, list2 *listnode.ListNode) *listnode.ListNode {
	result := listnode.ListNode{}

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
func mergeTwoLists2(list1 *listnode.ListNode, list2 *listnode.ListNode) *listnode.ListNode {
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
