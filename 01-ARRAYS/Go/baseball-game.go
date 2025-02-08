package main

import (
	"strconv"
)

func main() {
	operations := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	println(calPoints(operations))
}

func calPoints(operations []string) int {
	var stack [1000]int
	var stackI int
	for _, op := range operations {
		if op == "C" {
			stackI -= 1
			continue
		}

		var tmp int
		if op == "+" {
			tmp = stack[stackI-2] + stack[stackI-1]
		} else if op == "D" {
			tmp = stack[stackI-1] * 2
		} else {
			tmp, _ = strconv.Atoi(op)
		}

		stack[stackI] = tmp
		stackI++
	}

	var points int
	for _, j := range stack[0:stackI] {
		points += j
	}

	return points
}
