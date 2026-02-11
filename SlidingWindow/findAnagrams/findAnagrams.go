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
	if pl > sl {
		return []int{}
	}
	var need, win [26]int

	for i := 0; i < pl; i++ {
		need[p[i]-'a']++
		win[s[i]-'a']++
	}

	res := make([]int, 0)

	equal := func(a, b *[26]int) bool {
		for i := 0; i < 26; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	if equal(&need, &win) {
		res = append(res, 0)
	}

	for r := pl; r < sl; r++ {
		win[s[r]-'a']++
		win[s[r-pl]-'a']--
		if equal(&need, &win) {
			res = append(res, r-pl+1)
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
