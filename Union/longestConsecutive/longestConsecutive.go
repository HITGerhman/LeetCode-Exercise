package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	ans := 0
	for num := range set {
		if _, ok := set[num-1]; ok {
			continue
		}

		cur := num
		curLen := 1
		for {
			if _, ok := set[cur+1]; !ok {
				break
			}
			cur++
			curLen++
		}

		if curLen > ans {
			ans = curLen
		}
	}

	return ans
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	var nums []int
	_ = json.Unmarshal([]byte(in.Text()), &nums)
	fmt.Println(longestConsecutive(nums))
}
