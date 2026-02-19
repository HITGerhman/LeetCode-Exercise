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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == levelSize-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

func buildTree(arr []any) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: int(arr[0].(float64))}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		cur := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != nil {
			cur.Left = &TreeNode{Val: int(arr[i].(float64))}
			queue = append(queue, cur.Left)
		}
		i++
		if i < len(arr) && arr[i] != nil {
			cur.Right = &TreeNode{Val: int(arr[i].(float64))}
			queue = append(queue, cur.Right)
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
	fmt.Println(rightSideView(buildTree(arr)))
}
