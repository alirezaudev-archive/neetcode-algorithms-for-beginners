package main

// https://leetcode.com/problems/course-schedule/

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int, numCourses*2)
	for _, pair := range prerequisites {
		course := pair[0]
		graph[course] = append(graph[course], pair[1])
	}

	visited := make([]bool, numCourses)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if visited[node] {
			return false
		}
		if len(graph[node]) == 0 {
			return true
		}

		visited[node] = true

		for _, neighbour := range graph[node] {
			if !dfs(neighbour) {
				return false
			}
		}

		visited[node] = false
		graph[node] = []int{}
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
