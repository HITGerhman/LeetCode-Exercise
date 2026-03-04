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
		if start == n {
			if len(path) == 4 {
				res = append(res, strings.Join(path, "."))
			}
			return
		}
		remainChar := n - start
		remainSeg := 4 - len(path)
		if remainChar < remainSeg || remainChar > remainSeg*3 {
			return
		}
		for i := 1; i <= 3; i++ {
			if start+i > n {
				break
			}
			if i > 1 && s[start] == '0' {
				break
			}
			sub := s[start : start+i]
			if i == 3 && sub > "255" {
				break
			}
			path = append(path, sub)
			dfs(start + i)
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
