package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= r {
			if intervals[i][1] > r {
				r = intervals[i][1]
			}
		} else {
			res = append(res, []int{l, r})
			l, r = intervals[i][0], intervals[i][1]
		}
	}
	res = append(res, []int{l, r})
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)
	in.Scan()
	var intervals [][]int
	_ = json.Unmarshal([]byte(in.Text()), &intervals)
	fmt.Println(merge(intervals))
}
