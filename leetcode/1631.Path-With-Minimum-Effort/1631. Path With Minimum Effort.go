package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/template"
)

var dir = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 解法一 DFS + 二分
func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	low, high := 0, 1000000
	for low < high {
		threshold := low + (high-low)>>1
		if !hasPath(heights, visited, 0, 0, threshold) {
			low = threshold + 1
		} else {
			high = threshold
		}
		for i := range visited {
			for j := range visited[i] {
				visited[i][j] = false
			}
		}
	}
	return low
}

func hasPath(heights [][]int, visited [][]bool, i, j, threshold int) bool {
	n, m := len(heights), len(heights[0])
	if i == n-1 && j == m-1 {
		return true
	}
	visited[i][j] = true
	res := false
	for _, d := range dir {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || ni >= n || nj < 0 || nj >= m || visited[ni][nj] || res {
			continue
		}
		diff := abs(heights[i][j] - heights[ni][nj])
		if diff <= threshold && hasPath(heights, visited, ni, nj, threshold) {
			res = true
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 解法二 并查集
func minimumEffortPath1(heights [][]int) int {
	n, m, edges, uf := len(heights), len(heights[0]), []edge{}, template.UnionFind{}
	uf.Init(n * m)
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].diff < edges[j].diff })
	for _, e := range edges {
		uf.Union(e.v, e.w)
		if uf.Find(0) == uf.Find(n*m-1) {
			return e.diff
		}
	}
	return 0
}

type edge struct {
	v, w, diff int
}
