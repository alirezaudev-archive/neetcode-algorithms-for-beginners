package main

// https://leetcode.com/problems/clone-graph/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := map[*Node]*Node{}

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if v, ok := visited[node]; ok {
			return v
		}

		clone := &Node{Val: node.Val}
		visited[node] = clone
		for _, nei := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(nei))
		}

		return clone
	}

	return dfs(node)
}
