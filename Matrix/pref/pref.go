package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	pref := make([][]int, n+1)
	for i := range pref {
		pref[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		for j, c := range s {
			val := int(c - '0')
			pref[i+1][j+1] = val + pref[i][j+1] + pref[i+1][j] - pref[i][j]
		}
	}

	for k := 1; k <= n; k++ {
		if (k*k)%2 != 0 {
			fmt.Println(0)
			continue
		}

		count := 0
		target := (k * k) / 2

		for i := 0; i <= n-k; i++ {
			for j := 0; j <= n-k; j++ {
				sum := pref[i+1][j+1] - pref[i][j+1] - pref[i+1][j] - pref[i][j]
				if sum == target {
					count++
				}
			}
		}
		fmt.Println(count)
	}
}
