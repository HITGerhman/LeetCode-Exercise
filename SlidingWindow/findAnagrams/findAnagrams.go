package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findAnagrams(s string, p string) []int {
	sl := len(s)
	pl := len(p)

	if pl == 0 || pl > sl {
		return []int{}
	}

	need := make(map[byte]int)
	win := make(map[byte]int)

	for i := 0; i < pl; i++ {
		win[s[i]]++
		need[p[i]]++
	}

	equal := func(win map[byte]int, need map[byte]int) bool {
		for k, v := range need {
			if win[k] != v {
				return false
			}
		}
		return true
	}

	var res []int
	if equal(win, need) {
		res = append(res, 0)
	}
	for i := pl; i < sl; i++ {
		win[s[i]]++    // 右侧字符进窗口
		win[s[i-pl]]-- // 左侧字符出窗口
		if equal(win, need) {
			res = append(res, i-pl+1)
		}
	}
	return res
}

func main() {
	f, err := os.Open("input.txt")
	var in *bufio.Scanner
	if err != nil {
		in = bufio.NewScanner(os.Stdin)
	} else {
		defer f.Close()
		in = bufio.NewScanner(f)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	unquote := func(s string) string {
		s = strings.TrimSpace(s)
		if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
		return s
	}
	for {
		if !in.Scan() {
			break
		}
		s := unquote(in.Text())
		if s == "" {
			break
		}
		if !in.Scan() {
			break
		}
		p := unquote(in.Text())
		if p == "" {
			break
		}
		fmt.Fprintln(out, findAnagrams(s, p))
	}
}
