package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}

	res := make([]string, 0)
	path := make([]string, 0, 4)

	var dfs func(start int)
	dfs = func(start int) {
		segments := len(path)
		if segments == 4 {
			if start == n {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		remainChars := n - start
		remainSeg := 4 - segments
		if remainChars < remainSeg || remainChars > remainSeg*3 {
			return
		}

		num := 0
		for end := start; end < n && end < start+3; end++ {
			if end > start && s[start] == '0' {
				break
			}

			num = num*10 + int(s[end]-'0')
			if num > 255 {
				break
			}

			path = append(path, s[start:end+1])
			dfs(end + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	in := bufio.NewScanner(f)
	in.Scan()

	raw := in.Text()
	s := raw
	var parsed string
	if err := json.Unmarshal([]byte(raw), &parsed); err == nil {
		s = parsed
	}

	fmt.Println(restoreIpAddresses(s))
}
