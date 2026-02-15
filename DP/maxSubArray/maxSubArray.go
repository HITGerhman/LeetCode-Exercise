package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func maxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}
		if cur > best {
			best = cur
		}
	}
	return best
}
func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	var nums []int
	_ = json.Unmarshal([]byte(in.Text()), &nums)
	fmt.Println(maxSubArray(nums))
}
