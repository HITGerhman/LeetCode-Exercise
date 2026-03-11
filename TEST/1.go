package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// 使用二维前缀和数组 pref
	// pref[i][j] 表示矩阵 grid[0..i-1][0..j-1] 的元素之和
	pref := make([][]int, n+1)
	for i := range pref {
		pref[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		for j, c := range s {
			val := int(c - '0')
			// 计算前缀和: 当前值 + 上方 + 左方 - 左上方
			pref[i+1][j+1] = val + pref[i][j+1] + pref[i+1][j] - pref[i][j]
		}
	}

	// 枚举所有可能的正方形边长 k
	for k := 1; k <= n; k++ {
		// 如果总面积是奇数，不可能 0 和 1 数量相等，直接输出 0
		if (k*k)%2 != 0 {
			fmt.Println(0)
			continue
		}

		count := 0
		target := (k * k) / 2 // 目标是 1 的数量等于总面积的一半

		// 枚举左上角坐标 (i, j)
		for i := 0; i <= n-k; i++ {
			for j := 0; j <= n-k; j++ {
				// 利用前缀和 O(1) 计算子矩阵内 1 的数量
				// 子矩阵范围: 行 [i, i+k-1], 列 [j, j+k-1]
				// 对应前缀和坐标: (i+k, j+k), (i, j+k), (i+k, j), (i, j)
				sum := pref[i+k][j+k] - pref[i][j+k] - pref[i+k][j] + pref[i][j]
				if sum == target {
					count++
				}
			}
		}
		fmt.Println(count)
	}
}
