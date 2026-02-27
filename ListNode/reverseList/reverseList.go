package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // 1. 暂存后继节点
		curr.Next = prev  // 2. 修改引用指向前驱
		prev = curr       // 3. prev 前移
		curr = next       // 4. curr 前移
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	in := bufio.NewScanner(f)

	for in.Scan() {
		var nums []int
		if err := json.Unmarshal([]byte(in.Text()), &nums); err != nil {
			continue
		}
		head := intsToList(nums)
		newHead := reverseList(head)
		printList(newHead)
	}
}

func intsToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func printList(head *ListNode) {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	fmt.Println(res)
}
