//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k, a, b int
	fmt.Fscan(in, &n, &k, &a, &b)

	cap := make([]int, n)
	for i := range cap {
		fmt.Fscan(in, &cap[i])
	}

	cor := make([]int, n)
	for i := range cor {
		fmt.Fscan(in, &cor[i])
	}

	var ans int
	var validCnt int // 当前连续达标的人数

	for i := 0; i < n; i++ {
		// 如果这个人的各项数值都达标
		if cap[i] >= a && cor[i] >= b {
			validCnt++
		} else {
			validCnt = 0 // 一旦断开，重新计数
		}

		// 如果连续达标人数达到 k，说明以当前 i 结尾的这 k 个人是一个合法方案
		if validCnt >= k {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}
