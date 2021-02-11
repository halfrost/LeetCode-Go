---
title: 2.08 ✅ Backtracking
type: docs
weight: 8
---

# Backtracking

![](https://img.halfrost.com/Leetcode/Backtracking.png)

- 排列问题 Permutations。第 46 题，第 47 题。第 60 题，第 526 题，第 996 题。
- 组合问题 Combination。第 39 题，第 40 题，第 77 题，第 216 题。
- 排列和组合杂交问题。第 1079 题。
- N 皇后终极解法(二进制解法)。第 51 题，第 52 题。
- 数独问题。第 37 题。
- 四个方向搜索。第 79 题，第 212 题，第 980 题。
- 子集合问题。第 78 题，第 90 题。
- Trie。第 208 题，第 211 题。
- BFS 优化。第 126 题，第 127 题。
- DFS 模板。(只是一个例子，不对应任何题)

```go
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findcombinationSum2(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum2(nums []int, target, index int, c []int, res *[][]int) {
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] { // 这里是去重的关键逻辑
			continue
		}
		if target >= nums[i] {
			c = append(c, nums[i])
			findcombinationSum2(nums, target-nums[i], i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
```
- BFS 模板。(只是一个例子，不对应任何题)

```go
func updateMatrix_BFS(matrix [][]int) [][]int {
	res := make([][]int, len(matrix))
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	queue := make([][]int, 0)
	for i, _ := range matrix {
		res[i] = make([]int, len(matrix[0]))
		for j, _ := range res[i] {
			if matrix[i][j] == 0 {
				res[i][j] = -1
				queue = append(queue, []int{i, j})
			}
		}
	}
	level := 1
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size -= 1
			node := queue[0]
			queue = queue[1:]
			i, j := node[0], node[1]
			for _, direction := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
				x := i + direction[0]
				y := j + direction[1]
				if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || res[x][y] < 0 || res[x][y] > 0 {
					continue
				}
				res[x][y] = level
				queue = append(queue, []int{x, y})
			}
		}
		level++
	}
	for i, row := range res {
		for j, cell := range row {
			if cell == -1 {
				res[i][j] = 0
			}
		}
	}
	return res
}
```

| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0017|Letter Combinations of a Phone Number|[Go]({{< relref "/ChapterFour/0001~0099/0017.Letter-Combinations-of-a-Phone-Number.md" >}})|Medium| O(log n)| O(1)||49.0%|
|0022|Generate Parentheses|[Go]({{< relref "/ChapterFour/0001~0099/0022.Generate-Parentheses.md" >}})|Medium| O(log n)| O(1)||65.1%|
|0037|Sudoku Solver|[Go]({{< relref "/ChapterFour/0001~0099/0037.Sudoku-Solver.md" >}})|Hard| O(n^2)| O(n^2)|❤️|46.3%|
|0039|Combination Sum|[Go]({{< relref "/ChapterFour/0001~0099/0039.Combination-Sum.md" >}})|Medium| O(n log n)| O(n)||59.1%|
|0040|Combination Sum II|[Go]({{< relref "/ChapterFour/0001~0099/0040.Combination-Sum-II.md" >}})|Medium| O(n log n)| O(n)||50.0%|
|0046|Permutations|[Go]({{< relref "/ChapterFour/0001~0099/0046.Permutations.md" >}})|Medium| O(n)| O(n)|❤️|66.4%|
|0047|Permutations II|[Go]({{< relref "/ChapterFour/0001~0099/0047.Permutations-II.md" >}})|Medium| O(n^2)| O(n)|❤️|49.3%|
|0051|N-Queens|[Go]({{< relref "/ChapterFour/0001~0099/0051.N-Queens.md" >}})|Hard| O(n!)| O(n)|❤️|49.4%|
|0052|N-Queens II|[Go]({{< relref "/ChapterFour/0001~0099/0052.N-Queens-II.md" >}})|Hard| O(n!)| O(n)|❤️|60.0%|
|0060|Permutation Sequence|[Go]({{< relref "/ChapterFour/0001~0099/0060.Permutation-Sequence.md" >}})|Hard| O(n log n)| O(1)||39.3%|
|0077|Combinations|[Go]({{< relref "/ChapterFour/0001~0099/0077.Combinations.md" >}})|Medium| O(n)| O(n)|❤️|57.3%|
|0078|Subsets|[Go]({{< relref "/ChapterFour/0001~0099/0078.Subsets.md" >}})|Medium| O(n^2)| O(n)|❤️|64.9%|
|0079|Word Search|[Go]({{< relref "/ChapterFour/0001~0099/0079.Word-Search.md" >}})|Medium| O(n^2)| O(n^2)|❤️|36.7%|
|0089|Gray Code|[Go]({{< relref "/ChapterFour/0001~0099/0089.Gray-Code.md" >}})|Medium| O(n)| O(1)||50.3%|
|0090|Subsets II|[Go]({{< relref "/ChapterFour/0001~0099/0090.Subsets-II.md" >}})|Medium| O(n^2)| O(n)|❤️|48.7%|
|0093|Restore IP Addresses|[Go]({{< relref "/ChapterFour/0001~0099/0093.Restore-IP-Addresses.md" >}})|Medium| O(n)| O(n)|❤️|37.4%|
|0126|Word Ladder II|[Go]({{< relref "/ChapterFour/0100~0199/0126.Word-Ladder-II.md" >}})|Hard| O(n)| O(n^2)|❤️|23.6%|
|0131|Palindrome Partitioning|[Go]({{< relref "/ChapterFour/0100~0199/0131.Palindrome-Partitioning.md" >}})|Medium| O(n)| O(n^2)|❤️|51.8%|
|0211|Design Add and Search Words Data Structure|[Go]({{< relref "/ChapterFour/0200~0299/0211.Design-Add-and-Search-Words-Data-Structure.md" >}})|Medium| O(n)| O(n)|❤️|40.0%|
|0212|Word Search II|[Go]({{< relref "/ChapterFour/0200~0299/0212.Word-Search-II.md" >}})|Hard| O(n^2)| O(n^2)|❤️|36.9%|
|0216|Combination Sum III|[Go]({{< relref "/ChapterFour/0200~0299/0216.Combination-Sum-III.md" >}})|Medium| O(n)| O(1)|❤️|60.2%|
|0306|Additive Number|[Go]({{< relref "/ChapterFour/0300~0399/0306.Additive-Number.md" >}})|Medium| O(n^2)| O(1)|❤️|29.6%|
|0357|Count Numbers with Unique Digits|[Go]({{< relref "/ChapterFour/0300~0399/0357.Count-Numbers-with-Unique-Digits.md" >}})|Medium| O(1)| O(1)||48.8%|
|0401|Binary Watch|[Go]({{< relref "/ChapterFour/0400~0499/0401.Binary-Watch.md" >}})|Easy| O(1)| O(1)||48.4%|
|0526|Beautiful Arrangement|[Go]({{< relref "/ChapterFour/0500~0599/0526.Beautiful-Arrangement.md" >}})|Medium| O(n^2)| O(1)|❤️|61.8%|
|0784|Letter Case Permutation|[Go]({{< relref "/ChapterFour/0700~0799/0784.Letter-Case-Permutation.md" >}})|Medium| O(n)| O(n)||66.5%|
|0842|Split Array into Fibonacci Sequence|[Go]({{< relref "/ChapterFour/0800~0899/0842.Split-Array-into-Fibonacci-Sequence.md" >}})|Medium| O(n^2)| O(1)|❤️|36.8%|
|0980|Unique Paths III|[Go]({{< relref "/ChapterFour/0900~0999/0980.Unique-Paths-III.md" >}})|Hard| O(n log n)| O(n)||77.2%|
|0996|Number of Squareful Arrays|[Go]({{< relref "/ChapterFour/0900~0999/0996.Number-of-Squareful-Arrays.md" >}})|Hard| O(n log n)| O(n) ||48.3%|
|1079|Letter Tile Possibilities|[Go]({{< relref "/ChapterFour/1000~1099/1079.Letter-Tile-Possibilities.md" >}})|Medium| O(n^2)| O(1)|❤️|75.9%|
|1641|Count Sorted Vowel Strings|[Go]({{< relref "/ChapterFour/1600~1699/1641.Count-Sorted-Vowel-Strings.md" >}})|Medium||||77.0%|
|1655|Distribute Repeating Integers|[Go]({{< relref "/ChapterFour/1600~1699/1655.Distribute-Repeating-Integers.md" >}})|Hard||||40.4%|
|1659|Maximize Grid Happiness|[Go]({{< relref "/ChapterFour/1600~1699/1659.Maximize-Grid-Happiness.md" >}})|Hard||||35.4%|
|1681|Minimum Incompatibility|[Go]({{< relref "/ChapterFour/1600~1699/1681.Minimum-Incompatibility.md" >}})|Hard||||35.4%|
|1688|Count of Matches in Tournament|[Go]({{< relref "/ChapterFour/1600~1699/1688.Count-of-Matches-in-Tournament.md" >}})|Easy||||82.6%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Dynamic_Programming/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Depth_First_Search/">下一页➡️</a></p>
</div>
