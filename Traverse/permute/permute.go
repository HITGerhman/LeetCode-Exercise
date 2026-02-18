package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			one := make([]int, len(path))
			copy(one, path)
			res = append(res, one)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
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
