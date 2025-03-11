package main

import "fmt"

// https://leetcode.com/problems/subsets/

func main() {
	fmt.Printf("%v\n", subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	subs := [][]int{{}}

	for _, num := range nums {
		var tmp [][]int
		for _, sub := range subs {
			newSub := append([]int{}, sub...)
			tmp = append(tmp, append(newSub, num))
		}
		subs = append(subs, tmp...)
	}

	return subs
}
