package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j)
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	in := bufio.NewScanner(f)

	in.Scan()
	var grid [][]int
	_ = json.Unmarshal([]byte(in.Text()), &grid)
	fmt.Println(maxAreaOfIsland(grid))
}
