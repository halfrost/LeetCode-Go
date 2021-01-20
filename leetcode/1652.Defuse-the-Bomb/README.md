# [1652. Defuse the Bomb](https://leetcode.com/problems/defuse-the-bomb/)


## 题目

You have a bomb to defuse, and your time is running out! Your informer will provide you with a **circular** array `code` of length of `n` and a key `k`.

To decrypt the code, you must replace every number. All the numbers are replaced **simultaneously**.

- If `k > 0`, replace the `ith` number with the sum of the **next** `k` numbers.
- If `k < 0`, replace the `ith` number with the sum of the **previous** `k` numbers.
- If `k == 0`, replace the `ith` number with `0`.

As `code` is circular, the next element of `code[n-1]` is `code[0]`, and the previous element of `code[0]` is `code[n-1]`.

Given the **circular** array `code` and an integer key `k`, return *the decrypted code to defuse the bomb*!

**Example 1:**

```
Input: code = [5,7,1,4], k = 3
Output: [12,10,16,13]
Explanation: Each number is replaced by the sum of the next 3 numbers. The decrypted code is [7+1+4, 1+4+5, 4+5+7, 5+7+1]. Notice that the numbers wrap around.
```

**Example 2:**

```
Input: code = [1,2,3,4], k = 0
Output: [0,0,0,0]
Explanation: When k is zero, the numbers are replaced by 0. 
```

**Example 3:**

```
Input: code = [2,4,9,3], k = -2
Output: [12,5,6,13]
Explanation: The decrypted code is [3+9, 2+3, 4+2, 9+4]. Notice that the numbers wrap around again. If k is negative, the sum is of the previous numbers.
```

**Constraints:**

- `n == code.length`
- `1 <= n <= 100`
- `1 <= code[i] <= 100`
- `(n - 1) <= k <= n - 1`

## 题目大意

你有一个炸弹需要拆除，时间紧迫！你的情报员会给你一个长度为 n 的 循环 数组 code 以及一个密钥 k 。为了获得正确的密码，你需要替换掉每一个数字。所有数字会 同时 被替换。

- 如果 k > 0 ，将第 i 个数字用 接下来 k 个数字之和替换。
- 如果 k < 0 ，将第 i 个数字用 之前 k 个数字之和替换。
- 如果 k == 0 ，将第 i 个数字用 0 替换。

由于 code 是循环的， code[n-1] 下一个元素是 code[0] ，且 code[0] 前一个元素是 code[n-1] 。

给你 循环 数组 code 和整数密钥 k ，请你返回解密后的结果来拆除炸弹！

## 解题思路

- 给出一个 code 数组，要求按照规则替换每个字母。
- 简单题，按照题意描述循环即可。

## 代码

```go
package leetcode

func decrypt(code []int, k int) []int {
	if k == 0 {
		for i := 0; i < len(code); i++ {
			code[i] = 0
		}
		return code
	}
	count, sum, res := k, 0, make([]int, len(code))
	if k > 0 {
		for i := 0; i < len(code); i++ {
			for j := i + 1; j < len(code); j++ {
				if count == 0 {
					break
				}
				sum += code[j]
				count--
			}
			if count > 0 {
				for j := 0; j < len(code); j++ {
					if count == 0 {
						break
					}
					sum += code[j]
					count--
				}
			}
			res[i] = sum
			sum, count = 0, k
		}
	}
	if k < 0 {
		for i := 0; i < len(code); i++ {
			for j := i - 1; j >= 0; j-- {
				if count == 0 {
					break
				}
				sum += code[j]
				count++
			}
			if count < 0 {
				for j := len(code) - 1; j >= 0; j-- {
					if count == 0 {
						break
					}
					sum += code[j]
					count++
				}
			}
			res[i] = sum
			sum, count = 0, k
		}
	}
	return res
}
```