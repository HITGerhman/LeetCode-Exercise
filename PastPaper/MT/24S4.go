//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

type E struct{ u, v int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	initEdges := make([]E, m)
	initSet := map[E]bool{}
	all := make([]int, 0, 2*m+2*q)

}
