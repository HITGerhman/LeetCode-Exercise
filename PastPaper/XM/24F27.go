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

	var n, m int
	fmt.Fscan(in, &n, &m)

	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	ops := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &ops[i][0], &ops[i][1], &ops[i][2])
	}

	// 二分查找答案：范围 [1, m]
	l, r := 1, m
	ans := -1

	for l <= r {
		mid := (l + r) / 2
		if check(n, mid, nums, ops) {
			ans = mid
			r = mid - 1 // 尝试找更早的时刻
		} else {
			l = mid + 1 // 当前次数还不够，需要更多操作
		}
	}
	fmt.Fprint(out, ans)
}

func check(n, k int, nums []int, ops [][3]int) bool {
	// 使用差分数组优化区间减法，大小设为 n+2 以处理 1-based 索引和边界
	diff := make([]int, n+2)
	for i := 0; i < k; i++ {
		l, r, h := ops[i][0], ops[i][1], ops[i][2]
		diff[l] -= h
		diff[r+1] += h
	}

	currentDiff := 0
	for i := 1; i <= n; i++ {
		currentDiff += diff[i]
		// 检查：原高度 + 变动值 <= 0 ?
		// 注意：nums 是 0-based，而我们循环 i 是 1-based (对应 diff)
		if nums[i-1]+currentDiff <= 0 {
			return true
		}
	}
	return false
}
