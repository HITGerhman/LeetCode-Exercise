package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	ans := 1
	for i := 1; i < x; i++ {
		if x/i >= i {
			continue
		} else {
			ans = i - 1
			break
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
