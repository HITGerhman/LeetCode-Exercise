package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	cnt := make(map[byte]int)
	n := len(s)
	l, ans := 0, 0
	for i := 0; i < n; i++ {
		c := s[i]
		// 如果字符之前出现过，并且出现在当前窗口内（idx >= l），则移动左边界
		if idx, ok := cnt[c]; ok && idx >= l {
			l = idx + 1
		}
		cnt[c] = i
		// 更新最大长度
		if i-l+1 > ans {
			ans = i - l + 1
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
