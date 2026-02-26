package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var res []string
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}

	var path []string

	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == 4 {
			if start == n {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		// 剪枝：剩余字符数量必须在合理范围内（每段至少1个字符，至多3个）
		remainChar := n - start
		remainSeg := 4 - len(path)
		if remainChar < remainSeg || remainChar > remainSeg*3 {
			return
		}

		for i := 1; i <= 3; i++ {
			if start+i > n {
				break
			}
			sub := s[start : start+i]
			// 检查前导0：如果长度大于1且第一位是0，则非法
			if len(sub) > 1 && sub[0] == '0' {
				break
			}
			// 检查数值是否 <= 255
			val := 0
			for _, c := range sub {
				val = val*10 + int(c-'0')
			}
			if val > 255 {
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
