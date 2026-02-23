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
	ans := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		ans++
	}
	return ans
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
