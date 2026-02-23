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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	q := []*TreeNode{root}
	isToRight := true
	for len(q) > 0 {
		n := len(q)
		level := make([]int, n)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			idx := i
			if !isToRight {
				idx = n - 1 - i
			}
			level[idx] = cur.Val
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, level)
		isToRight = !isToRight
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
	fmt.Println(zigzagLevelOrder(buildTree(arr)))
}
