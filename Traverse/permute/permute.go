package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func permute(nums []int) [][]int {
	res := [][]int{}
	level := []int{}
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(level) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, level)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				level = append(level, nums[i])
				dfs()
				level = level[:len(level)-1]
				used[i] = false
			}
		}
	}
	dfs()
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	var nums []int
	_ = json.Unmarshal([]byte(in.Text()), &nums)
	fmt.Println(permute(nums))
}
