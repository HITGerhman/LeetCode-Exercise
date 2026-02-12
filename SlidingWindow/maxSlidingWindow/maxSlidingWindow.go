package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k <= 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	res := make([]int, 0, n-k+1)
	deque := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	var nums []int
	_ = json.Unmarshal([]byte(in.Text()), &nums)
	in.Scan()
	k, _ := strconv.Atoi(in.Text())
	fmt.Println(maxSlidingWindow(nums, k))
}
