package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] > leftMax {
				leftMax = height[l]
			}
			ans += leftMax - height[l]
			l++
		} else {
			if height[r] > rightMax {
				rightMax = height[r]
			}
			ans += rightMax - height[r]
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
