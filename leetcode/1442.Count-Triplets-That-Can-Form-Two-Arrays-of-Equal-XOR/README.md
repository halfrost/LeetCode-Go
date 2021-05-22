# [1442. Count Triplets That Can Form Two Arrays of Equal XOR](https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/)


## 题目

Given an array of integers `arr`.

We want to select three indices `i`, `j` and `k` where `(0 <= i < j <= k < arr.length)`.

Let's define `a` and `b` as follows:

- `a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]`
- `b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]`

Note that **^** denotes the **bitwise-xor** operation.

Return *the number of triplets* (`i`, `j` and `k`) Where `a == b`.

**Example 1:**

```
Input: arr = [2,3,1,6,7]
Output: 4
Explanation: The triplets are (0,1,2), (0,2,2), (2,3,4) and (2,4,4)
```

**Example 2:**

```
Input: arr = [1,1,1,1,1]
Output: 10
```

**Example 3:**

```
Input: arr = [2,3]
Output: 0
```

**Example 4:**

```
Input: arr = [1,3,5,7,9]
Output: 3
```

**Example 5:**

```
Input: arr = [7,11,12,9,5,2,7,17,22]
Output: 8
```

**Constraints:**

- `1 <= arr.length <= 300`
- `1 <= arr[i] <= 10^8`

## 题目大意

给你一个整数数组 arr 。现需要从数组中取三个下标 i、j 和 k ，其中 (0 <= i < j <= k < arr.length) 。a 和 b 定义如下：

- a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
- b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]

注意：^ 表示 按位异或 操作。请返回能够令 a == b 成立的三元组 (i, j , k) 的数目。

## 解题思路

- 这一题需要用到 `x^x = 0` 这个异或特性。题目要求 `a == b`，可以等效转化为 `arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1] ^ arr[j] ^ arr[j + 1] ^ ... ^ arr[k] = 0`，这样 j 相当于可以“忽略”，专注找到所有元素异或结果为 0 的区间 [i,k] 即为答案。利用前缀和的思想，只不过此题非累加和，而是异或。又由 `x^x = 0` 这个异或特性，相同部分异或相当于消除，于是有 `prefix[i,k] = prefix[0,k] ^ prefix[0,i-1]`，找到每一个 `prefix[i,k] = 0` 的 i，k 组合，i < j <= k，那么满足条件的三元组 (i,j,k) 的个数完全取决于 j 的取值范围，(因为 i 和 k 已经固定了)，j 的取值范围为 k-i，所以累加所有满足条件的 k-i，输出即为最终答案。

## 代码

```go
package leetcode

func countTriplets(arr []int) int {
	prefix, num, count, total := make([]int, len(arr)), 0, 0, 0
	for i, v := range arr {
		num ^= v
		prefix[i] = num
	}
	for i := 0; i < len(prefix)-1; i++ {
		for k := i + 1; k < len(prefix); k++ {
			total = prefix[k]
			if i > 0 {
				total ^= prefix[i-1]
			}
			if total == 0 {
				count += k - i
			}
		}
	}
	return count
}
```