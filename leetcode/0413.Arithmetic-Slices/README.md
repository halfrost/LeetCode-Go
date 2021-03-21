# [413. Arithmetic Slices](https://leetcode.com/problems/arithmetic-slices/)


## 题目

A sequence of numbers is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.

For example, these are arithmetic sequences:

```
1, 3, 5, 7, 9
7, 7, 7, 7
3, -1, -5, -9
```

The following sequence is not arithmetic.

```
1, 1, 2, 5, 7
```

A zero-indexed array A consisting of N numbers is given. A slice of that array is any pair of integers (P, Q) such that 0 <= P < Q < N.

A slice (P, Q) of the array A is called arithmetic if the sequence:A[P], A[P + 1], ..., A[Q - 1], A[Q] is arithmetic. In particular, this means that P + 1 < Q.

The function should return the number of arithmetic slices in the array A.

**Example:**

```
A = [1, 2, 3, 4]

return: 3, for 3 arithmetic slices in A: [1, 2, 3], [2, 3, 4] and [1, 2, 3, 4] itself.
```

## 题目大意

数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N 。如果满足以下条件，则称子数组(P, Q)为等差数组：元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。函数要返回数组 A 中所有为等差数组的子数组个数。

## 解题思路

- 由题目给出的定义，至少 3 个数字以上的等差数列才满足题意。连续 k 个连续等差的元素，包含的子等差数列是底层的，1，2，3…… k。所以每判断一组 3 个连续的数列，只需要用一个变量累加前面已经有多少个满足题意的连续元素，只要满足题意的等差数列就加上这个累加值。一旦不满足等差的条件，累加值置 0。如此循环一次即可找到题目要求的答案。

## 代码

```go
package leetcode

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}
	res, dp := 0, 0
	for i := 1; i < len(A)-1; i++ {
		if A[i+1]-A[i] == A[i]-A[i-1] {
			dp++
			res += dp
		} else {
			dp = 0
		}
	}
	return res
}
```