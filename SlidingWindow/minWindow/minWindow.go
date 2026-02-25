package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func minWindow(s string, t string) string {
	sl := len(s)
	tl := len(t)

	if tl == 0 || sl < tl {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < tl; i++ {
		need[t[i]]++
	}
	win := make(map[byte]int)
	valid := 0
	kindsNeed := len(need)
	bestL, bestLen := 0, sl+1
	l := 0
	for i := 0; i < sl; i++ {
		c := s[i]
		if _, ok := need[c]; ok {
			win[c]++
			if win[c] == need[c] {
				valid++
			}
		}
		for valid == kindsNeed {
			if i-l+1 < bestLen {
				bestL = l
				bestLen = i - l + 1
			}
			d := s[l]
			l++
			if _, ok := need[d]; ok {
				if win[d] == need[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	if bestLen == sl+1 {
		return ""
	}
	return s[bestL : bestL+bestLen]
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
		t := unquote(in.Text())
		if t == "" {
			break
		}
		fmt.Fprintln(out, minWindow(s, t))
	}
}
