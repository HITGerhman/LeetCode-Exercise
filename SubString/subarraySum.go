package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func subarraySum(nums []int, k int) int {
	cnt:=make([]int,len(nums))
	cnt[0]=1
	
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for {
		if !in.Scan() {
			break
		}
		line := strings.TrimSpace(in.Text())
		line = line[1 : len(line)-1]
		parts := strings.Split(line, ",")
		nums := make([]int, 0, len(parts))
		if line != "" {
			for _, p := range parts {
				v, _ := strconv.Atoi(strings.TrimSpace(p))
				nums = append(nums, v)
			}
		}
		in.Scan()
		k, _ := strconv.Atoi(strings.TrimSpace(in.Text()))
		fmt.Fprintln(out, subarraySum(nums, k))
	}
}
