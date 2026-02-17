package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func search(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		}
		if nums[mid] > target {
			r = mid - 1
		}
	}
	return -1
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
	fmt.Println(search(nums, k))
}
