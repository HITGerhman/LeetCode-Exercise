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

	var n, x int
	fmt.Fscan(in, &n, &x)

	// 动态规划求解 0/1 背包问题
	// 注意：这个解法只有在背包容量 x 比较小的时候才可行 (例如 x <= 10^5)。
	// 对于 x 高达 10^9 的情况，这个方法会超内存且超时。
	dp := make([]int, x+1)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)

		// 从后往前遍历，确保每个物品只被使用一次（空间优化技巧）
		for j := x; j >= a; j-- {
			// 决策：不装入物品 (价值为 dp[j]) vs 装入物品 (价值为 dp[j-a] + b)
			if dp[j] < dp[j-a]+b {
				dp[j] = dp[j-a] + b
			}
		}
	}
	fmt.Fprintln(out, dp[x])
}
