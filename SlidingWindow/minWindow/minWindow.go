package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	win := make(map[byte]int)
	needKinds := len(need)
	valid := 0
	l := 0
	bestL, bestLen := 0, len(s)+1

	for r := 0; r < len(s); r++ {
		c := s[r]
		if _, ok := need[c]; ok {
			win[c]++
			if win[c] == need[c] {
				valid++
			}
		}

		for valid == needKinds {
			if r-l+1 < bestLen {
				bestLen = r - l + 1
				bestL = l
			}
			d := s[l]
			if _, ok := need[d]; ok {
				if win[d] == need[d] {
					valid--
				}
				win[d]--
			}
			l++
		}
	}

	if bestLen == len(s)+1 {
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
