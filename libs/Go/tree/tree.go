package tree

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func FromArray(tree []int) (root *Node) {
	if len(tree) == 0 {
		return &Node{Val: -1}
	}

	root = &Node{Val: tree[0]}
	layer := []*Node{root}
	i := 1
	for i < len(tree) {
		var tempLayer []*Node
		for _, node := range layer {
			if i < len(tree) && tree[i] != -1 {
				node.Left = &Node{Val: tree[i]}
				tempLayer = append(tempLayer, node.Left)
			}
			i++
			if i < len(tree) && tree[i] != -1 {
				node.Right = &Node{Val: tree[i]}
				tempLayer = append(tempLayer, node.Right)
			}
			i++
		}
		layer = tempLayer
	}
	return
}

func (this *Node) Print() *Node {
	height := this.Height()
	col := getCol(height)
	M := make([][]int, height)
	for i := range M {
		tmp := make([]int, col)
		for j := range tmp {
			tmp[j] = -1
		}
		M[i] = tmp
	}
	this.printTree(M, this, col/2, 0, height)
	for _, i := range M {
		for _, j := range i {
			if j < 0 {
				fmt.Print("..")
			} else {
				fmt.Print(strconv.Itoa(j) + ".")
			}
		}
		fmt.Println()
	}

	return this
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return max(n.Left.Height(), n.Right.Height()) + 1
}

func getCol(height int) int {
	return (1 << height) - 1
}

func (this *Node) printTree(M [][]int, root *Node, col, row, height int) {
	if root == nil || row >= len(M) || col < 0 || col >= len(M[0]) {
		return
	}

	M[row][col] = root.Val

	if height > 1 {
		shift := 1 << (height - 2)
		this.printTree(M, root.Left, col-shift, row+1, height-1)
		this.printTree(M, root.Right, col+shift, row+1, height-1)
	}
}
