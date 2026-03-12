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
	fmt.Fscan(in, &n, &k)

	// 预处理每个数字包含的因子 2 和 5 的个数
	cnt2 := make([]int, n)
	cnt5 := make([]int, n)
	var total2, total5 int

	for i := 0; i < n; i++ {
		var val int
		fmt.Fscan(in, &val)

		// 计算当前数字包含多少个因子 2 和 5
		for val > 0 && val%2 == 0 {
			cnt2[i]++
			val /= 2
		}
		for val > 0 && val%5 == 0 {
			cnt5[i]++
			val /= 5
		}
		total2 += cnt2[i]
		total5 += cnt5[i]
	}

	// 计算允许删除的因子最大数量
	// 如果剩余要 >= k，那么删除的就必须 <= 总数 - k
	limit2 := total2 - k
	limit5 := total5 - k

	if limit2 < 0 || limit5 < 0 {
		fmt.Fprintln(out, 0)
		return
	}

	var ans int
	var curr2, curr5 int
	left := 0
	// 滑动窗口：寻找满足 sum(区间) <= limit 的所有区间
	for right := 0; right < n; right++ {
		curr2 += cnt2[right]
		curr5 += cnt5[right]
		for curr2 > limit2 || curr5 > limit5 {
			curr2 -= cnt2[left]
			curr5 -= cnt5[left]
			left++
		}
		// 以 right 结尾的合法区间有 (right - left + 1) 个
		ans += right - left + 1
	}
	fmt.Fprintln(out, ans)
}

//amazing! Genius
//I got it!
