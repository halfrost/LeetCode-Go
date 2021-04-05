# [775. Global and Local Inversions](https://leetcode.com/problems/global-and-local-inversions/)


## 题目

We have some permutation `A` of `[0, 1, ..., N - 1]`, where `N` is the length of `A`.

The number of (global) inversions is the number of `i < j` with `0 <= i < j < N` and `A[i] > A[j]`.

The number of local inversions is the number of `i` with `0 <= i < N` and `A[i] > A[i+1]`.

Return `true` if and only if the number of global inversions is equal to the number of local inversions.

**Example 1:**

```
Input: A = [1,0,2]
Output: true
Explanation: There is 1 global inversion, and 1 local inversion.
```

**Example 2:**

```
Input: A = [1,2,0]
Output: false
Explanation: There are 2 global inversions, and 1 local inversion.
```

**Note:**

- `A` will be a permutation of `[0, 1, ..., A.length - 1]`.
- `A` will have length in range `[1, 5000]`.
- The time limit for this problem has been reduced.

## 题目大意

数组 A 是 [0, 1, ..., N - 1] 的一种排列，N 是数组 A 的长度。全局倒置指的是 i,j 满足 0 <= i < j < N 并且 A[i] > A[j] ，局部倒置指的是 i 满足 0 <= i < N 并且 A[i] > A[i+1] 。当数组 A 中全局倒置的数量等于局部倒置的数量时，返回 true 。

## 解题思路

- 本题代码非常简单，重在思考的过程。`[0, 1, ..., N - 1]` 不出现全局倒置的理想情况应该是 `i` 排列在 `A[i-1]`，`A[i]`，`A[i+1]` 的位置上。例如 `1` 如果排列在 `A[3]` 的位置上，那么比 `1` 小的只有 `0` 一个元素，`A[0]`，`A[1]`，`A[2]` 中必定有 2 个元素比 `1` 大，那必须会出现全局倒置的情况。`[0, 1, ..., N - 1]` 这是最理想的情况，每个元素都在自己的位置上。每个元素如果往左右相互偏移 1 个元素，那么也能保证只存在局部倒置，如果左右偏移 2 个元素，那必定会出现全局倒置。所以结论是：**不出现全局倒置的理想情况应该是 `i` 排列在 `A[i-1]`，`A[i]`，`A[i+1]` 的位置上**。判断这个结论的代码很简单，只需要判断 `A[i] - i` 的取值是否是 -1，0，1，也即 `abs(A[i] - i ) ≤ 1`。

## 代码

```go
package leetcode

func isIdealPermutation(A []int) bool {
	for i := range A {
		if abs(A[i]-i) > 1 {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```