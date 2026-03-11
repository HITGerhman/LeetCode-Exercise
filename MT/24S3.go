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

	var n, k int
	var s string
	fmt.Fscan(in, &n, &k)
	fmt.Fscan(in, &s)

	var ans int
	// 注意：range 字符串时，第一个返回值是下标，第二个才是字符
	for _, c := range s {
		if c == 'M' || c == 'T' {
			ans++
		}
	}

	// 计算结果，但最大不能超过字符串长度 n
	res := ans + k
	if res > n {
		res = n
	}
	fmt.Fprintln(out, res)
}
