package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	last := make([]int, 256)
	for i := range last {
		last[i] = -1
	}
	l, ans := 0, 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		if p := last[c]; p >= l {
			l = p + 1
		}
		last[c] = r
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for in.Scan() {
		line := strings.TrimSpace(in.Text())
		if line == "" {
			break
		}
		if len(line) >= 2 && line[0] == '"' && line[len(line)-1] == '"' {
			line = line[1 : len(line)-1]
		}
		fmt.Fprintln(out, lengthOfLongestSubstring(line))
	}
}
