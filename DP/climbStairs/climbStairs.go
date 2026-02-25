package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// 申请 n+1 的空间，以便直接使用 dp[n]
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	k, _ := strconv.Atoi(in.Text())
	fmt.Println(climbStairs(k))
}
