# [989. Add to Array-Form of Integer](https://leetcode.com/problems/add-to-array-form-of-integer/)

## 题目

For a non-negative integer `X`, the *array-form of `X`* is an array of its digits in left to right order.  For example, if `X = 1231`, then the array form is `[1,2,3,1]`.

Given the array-form `A` of a non-negative integer `X`, return the array-form of the integer `X+K`.

**Example 1:**

```
Input: A = [1,2,0,0], K = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234
```

**Example 2:**

```
Input: A = [2,7,4], K = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455
```

**Example 3:**

```
Input: A = [2,1,5], K = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021
```

**Example 4:**

```
Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
Output: [1,0,0,0,0,0,0,0,0,0,0]
Explanation: 9999999999 + 1 = 10000000000
```

**Note：**

1. `1 <= A.length <= 10000`
2. `0 <= A[i] <= 9`
3. `0 <= K <= 10000`
4. If `A.length > 1`, then `A[0] != 0`

## 题目大意

对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。

## 解题思路

- 简单题，计算 2 个非负整数的和。累加过程中不断的进位，最终输出到数组中记得需要逆序，即数字的高位排在数组下标较小的位置。

## 代码

```go
package leetcode

func addToArrayForm(A []int, K int) []int {
	res := []int{}
	for i := len(A) - 1; i >= 0; i-- {
		sum := A[i] + K%10
		K /= 10
		if sum >= 10 {
			K++
			sum -= 10
		}
		res = append(res, sum)
	}
	for ; K > 0; K /= 10 {
		res = append(res, K%10)
	}
	reverse(res)
	return res
}

func reverse(A []int) {
	for i, n := 0, len(A); i < n/2; i++ {
		A[i], A[n-1-i] = A[n-1-i], A[i]
	}
}
```