package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	maxLeft, maxRight := 0, 0
	ans := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] >= maxLeft {
				maxLeft = height[l]
			} else {
				ans += maxLeft - height[l]
			}
			l++
		} else {
			if height[r] >= maxRight {
				maxRight = height[r]
			} else {
				ans += maxRight - height[r]
			}
			r--
		}
	}
	return ans
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)
	in.Scan()
	var height []int
	_ = json.Unmarshal([]byte(in.Text()), &height)
	fmt.Println(trap(height))
}
