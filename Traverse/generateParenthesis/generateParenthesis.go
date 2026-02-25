package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(l, r int, path []byte)
	dfs = func(l, r int, path []byte) {
		if l == n && r == n {
			res = append(res, string(path))
			return
		}
		if l < n {
			path = append(path, '(')
			dfs(l+1, r, path)
			path = path[:len(path)-1] // 回溯：移除刚才添加的字符
		}
		if r < l {
			path = append(path, ')')
			dfs(l, r+1, path)
			path = path[:len(path)-1] // 回溯
		}
	}
	// 预分配容量为 2*n 的切片，避免 append 时扩容
	dfs(0, 0, make([]byte, 0, 2*n))
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
