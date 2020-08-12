# [598. Range Addition II](https://leetcode.com/problems/range-addition-ii/)


## 题目

Given an m * n matrix **M** initialized with all **0**'s and several update operations.

Operations are represented by a 2D array, and each operation is represented by an array with two **positive** integers **a** and **b**, which means **M[i][j]** should be **added by one** for all **0 <= i < a** and **0 <= j < b**.

You need to count and return the number of maximum integers in the matrix after performing all the operations.

**Example 1**:

```
Input: 
m = 3, n = 3
operations = [[2,2],[3,3]]
Output: 4
Explanation: 
Initially, M = 
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]

After performing [2,2], M = 
[[1, 1, 0],
 [1, 1, 0],
 [0, 0, 0]]

After performing [3,3], M = 
[[2, 2, 1],
 [2, 2, 1],
 [1, 1, 1]]

So the maximum integer in M is 2, and there are four of it in M. So return 4.
```

**Note**:

1. The range of m and n is [1,40000].
2. The range of a is [1,m], and the range of b is [1,n].
3. The range of operations size won't exceed 10,000.

## 题目大意

给定一个初始元素全部为 0，大小为 m*n 的矩阵 M 以及在 M 上的一系列更新操作。操作用二维数组表示，其中的每个操作用一个含有两个正整数 a 和 b 的数组表示，含义是将所有符合 0 <= i < a 以及 0 <= j < b 的元素 M[i][j] 的值都增加 1。在执行给定的一系列操作后，你需要返回矩阵中含有最大整数的元素个数。

注意:

- m 和 n 的范围是 [1,40000]。
- a 的范围是 [1,m]，b 的范围是 [1,n]。
- 操作数目不超过 10000。


## 解题思路

- 给定一个初始都为 0 的 m * n 的矩阵，和一个操作数组。经过一系列的操作以后，最终输出矩阵中最大整数的元素个数。每次操作都使得一个矩形内的元素都 + 1 。
- 这一题乍一看像线段树的区间覆盖问题，但是实际上很简单。如果此题是任意的矩阵，那就可能用到线段树了。这一题每个矩阵的起点都包含 [0 , 0] 这个元素，也就是说每次操作都会影响第一个元素。那么这道题就很简单了。经过 n 次操作以后，被覆盖次数最多的矩形区间，一定就是最大整数所在的区间。由于起点都是第一个元素，所以我们只用关心矩形的右下角那个坐标。右下角怎么计算呢？只用每次动态的维护一下矩阵长和宽的最小值即可。

## 代码

```go

package leetcode

func maxCount(m int, n int, ops [][]int) int {
	minM, minN := m, n
	for _, op := range ops {
		minM = min(minM, op[0])
		minN = min(minN, op[1])
	}
	return minM * minN
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```