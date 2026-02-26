package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	level := make([]int, 0)
	n := len(nums)

	var dfs func(start int)
	dfs = func(start int) {
		temp := make([]int, len(level))
		copy(temp, level)
		res = append(res, temp)
		for i := start; i < n; i++ {
			level = append(level, nums[i])
			dfs(i + 1)
			level = level[:len(level)-1]
		}
	}
	dfs(0)
	return res
}
func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)
	in.Scan()
	var nums []int
	_ = json.Unmarshal([]byte(in.Text()), &nums)
	fmt.Println(subsets(nums))
}
