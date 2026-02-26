package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	var dfs func(l, r int)
	dfs = func(l, r int) {
		if l == n && r == n {
			res = append(res, string(path))
			return
		}
		if l < n {
			path = append(path, '(')
			dfs(l+1, r)
			path = path[:len(path)-1]
		}
		if r < l {
			path = append(path, ')')
			dfs(l, r+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	fmt.Println(generateParenthesis(n))
}
