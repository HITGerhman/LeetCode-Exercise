package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}

func buildTree(arr []any) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}
	root := &TreeNode{Val: int(arr[0].(float64))}
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 && i < len(arr) {
		cur := q[0]
		q = q[1:]
		if i < len(arr) && arr[i] != nil {
			cur.Left = &TreeNode{Val: int(arr[i].(float64))}
			q = append(q, cur.Left)
		}
		i++
		if i < len(arr) && arr[i] != nil {
			cur.Right = &TreeNode{Val: int(arr[i].(float64))}
			q = append(q, cur.Right)
		}
		i++
	}
	return root
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)
	in.Scan()
	var arr []any
	_ = json.Unmarshal([]byte(in.Text()), &arr)
	fmt.Println(maxDepth(buildTree(arr)))
}
