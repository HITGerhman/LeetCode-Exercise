package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(l, r int, path string)
	dfs = func(l, r int, path string) {
		if l == n && r == n {
			res = append(res, path)
			return
		}
		if l < n {
			dfs(l+1, r, path+"(")
		}
		if r < l {
			dfs(l, r+1, path+")")
		}
	}
	dfs(0, 0, "")
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	fmt.Println(generateParenthesis(n))
}
