# [823. Binary Trees With Factors](https://leetcode.com/problems/binary-trees-with-factors/)


## 题目

Given an array of unique integers, `arr`, where each integer `arr[i]` is strictly greater than `1`.

We make a binary tree using these integers, and each number may be used for any number of times. Each non-leaf node's value should be equal to the product of the values of its children.

Return *the number of binary trees we can make*. The answer may be too large so return the answer **modulo** `109 + 7`.

**Example 1:**

```
Input: arr = [2,4]
Output: 3
Explanation: We can make these trees: [2], [4], [4, 2, 2]
```

**Example 2:**

```
Input: arr = [2,4,5,10]
Output: 7
Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
```

**Constraints:**

- `1 <= arr.length <= 1000`
- `2 <= arr[i] <= 10^9`

## 题目大意

给出一个含有不重复整数元素的数组，每个整数均大于 1。我们用这些整数来构建二叉树，每个整数可以使用任意次数。其中：每个非叶结点的值应等于它的两个子结点的值的乘积。满足条件的二叉树一共有多少个？返回的结果应模除 10 * 9 + 7。

## 解题思路

- 首先想到的是暴力解法，先排序，然后遍历所有节点，枚举两两乘积为第三个节点值的组合。然后枚举这些组合并构成树。这里计数的时候要注意，左右孩子如果不是对称的，左右子树相互对调又是一组解。但是这个方法超时了。原因是，暴力枚举了很多次重复的节点和组合。优化这里的方法就是把已经计算过的节点放入 `map` 中。这里有 2 层 `map`，第一层 `map` 记忆化的是两两乘积的组合，将父亲节点作为 `key`，左右 2 个孩子作为 `value`。第二层 `map` 记忆化的是以 `root` 为根节点此时二叉树的种类数，`key` 是 `root`，`value` 存的是种类数。这样优化以后，DFS 暴力解法可以 runtime beats 100%。
- 另外一种解法是 DP。定义 `dp[i]` 代表以 `i` 为根节点的树的种类数。dp[i] 初始都是 1，因为所有节点自身可以形成为自身单个节点为 `root` 的树。同样需要先排序。状态转移方程是：

    $$dp[i] = \sum_{j<i, k<i}^{}dp[j] * dp[k], j * k = i$$

    最后将 `dp[]` 数组中所有结果累加取模即为最终结果，时间复杂度 O(n^2)，空间复杂度 O(n)。

## 代码

```go
package leetcode

import (
	"sort"
)

const mod = 1e9 + 7

// 解法一 DFS
func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	numDict := map[int]bool{}
	for _, num := range arr {
		numDict[num] = true
	}
	dict, res := make(map[int][][2]int), 0
	for i, num := range arr {
		for j := i; j < len(arr) && num*arr[j] <= arr[len(arr)-1]; j++ {
			tmp := num * arr[j]
			if !numDict[tmp] {
				continue
			}
			dict[tmp] = append(dict[tmp], [2]int{num, arr[j]})
		}
	}
	cache := make(map[int]int)
	for _, num := range arr {
		res = (res + dfs(num, dict, cache)) % mod
	}
	return res
}

func dfs(num int, dict map[int][][2]int, cache map[int]int) int {
	if val, ok := cache[num]; ok {
		return val
	}
	res := 1
	for _, tuple := range dict[num] {
		a, b := tuple[0], tuple[1]
		x, y := dfs(a, dict, cache), dfs(b, dict, cache)
		tmp := x * y
		if a != b {
			tmp *= 2
		}
		res = (res + tmp) % mod
	}
	cache[num] = res
	return res
}

// 解法二 DP
func numFactoredBinaryTrees1(arr []int) int {
	dp := make(map[int]int)
	sort.Ints(arr)
	for i, curNum := range arr {
		for j := 0; j < i; j++ {
			factor := arr[j]
			quotient, remainder := curNum/factor, curNum%factor
			if remainder == 0 {
				dp[curNum] += dp[factor] * dp[quotient]
			}
		}
		dp[curNum]++
	}
	totalCount := 0
	for _, count := range dp {
		totalCount += count
	}
	return totalCount % mod
}
```