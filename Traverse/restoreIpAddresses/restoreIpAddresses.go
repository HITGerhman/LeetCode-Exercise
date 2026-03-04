package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res:=[]string{}
	n:=len(s)
	path:=make([]string,0,4)
	if n<4||n>12 {
		return []string{}
	}
	var dfs func(start int)
	dfs=func(start int) {
		if start==n {
			if len(path)==4 {
				res = append(res, strings.Join(path, "."))
			}
		}
		remainChar:=n-start
		remainSeg:=n-len(path)

		if remainChar<remainSeg||remainChar>remainSeg*3 {
			return
		}
		val:=0
		sub:=s[start:start+3]
		if sub[0]=='0' {
			return 
		}
		for i:=0;i<3;i++{
			if 
		}
	}
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
