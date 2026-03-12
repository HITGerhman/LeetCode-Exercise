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

	var n, q int
	fmt.Fscan(in, &n, &q)

	var baseSum int
	var zeroCount int

	for i := 0; i < n; i++ {
		var val int
		// Fscan 会自动跳过空格和换行，直接读取下一个整数
		fmt.Fscan(in, &val)
		if val == 0 {
			zeroCount++
		} else {
			baseSum += val
		}
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)

		// l 是允许填入的最小值，所以全部填 l 能得到最小和
		minSum := baseSum + zeroCount*l
		// r 是允许填入的最大值，所以全部填 r 能得到最大和
		maxSum := baseSum + zeroCount*r

		fmt.Fprintln(out, minSum, maxSum)
	}
}
