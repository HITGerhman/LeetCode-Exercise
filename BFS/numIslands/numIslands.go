package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func numIslands(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	ans := 0
	var bfs func(i int, j int)
	bfs = func(i int, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		bfs(i+1, j)
		bfs(i-1, j)
		bfs(i, j+1)
		bfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
				ans++
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
	fmt.Println(numIslands(grid))
}
