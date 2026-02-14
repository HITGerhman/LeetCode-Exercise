package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	l, r := 1, x/2
	ans := 0
	for l <= r {
		mid := l + (r-l)/2
		if mid <= x/mid {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	k, _ := strconv.Atoi(in.Text())
	fmt.Println(mySqrt(k))
}
