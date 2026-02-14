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
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	k, _ := strconv.Atoi(in.Text())
	fmt.Println(climbStairs(k))
}
