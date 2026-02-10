package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	last := make(map[byte]int)
	l, ans := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if p, ok := last[c]; ok && p > l {
			l = p
		}
		if i-l+1 > ans {
			ans = i - l + 1
		}
		last[c] = i + 1
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
