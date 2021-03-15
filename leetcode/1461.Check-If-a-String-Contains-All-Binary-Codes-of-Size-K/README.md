# [1461. Check If a String Contains All Binary Codes of Size K](https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/)


## 题目

Given a binary string `s` and an integer `k`.

Return *True* if every binary code of length `k` is a substring of `s`. Otherwise, return *False*.

**Example 1:**

```
Input: s = "00110110", k = 2
Output: true
Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They can be all found as substrings at indicies 0, 1, 3 and 2 respectively.
```

**Example 2:**

```
Input: s = "00110", k = 2
Output: true
```

**Example 3:**

```
Input: s = "0110", k = 1
Output: true
Explanation: The binary codes of length 1 are "0" and "1", it is clear that both exist as a substring. 
```

**Example 4:**

```
Input: s = "0110", k = 2
Output: false
Explanation: The binary code "00" is of length 2 and doesn't exist in the array.
```

**Example 5:**

```
Input: s = "0000000001011100", k = 4
Output: false
```

**Constraints:**

- `1 <= s.length <= 5 * 10^5`
- `s` consists of 0's and 1's only.
- `1 <= k <= 20`

## 题目大意

给你一个二进制字符串 `s` 和一个整数 `k` 。如果所有长度为 `k` 的二进制字符串都是 `s` 的子串，请返回 `True` ，否则请返回 `False` 。

## 解题思路

- 构造一个 `mask` 遮罩，依次划过整个二进制字符串，每次滑动即取出遮罩遮住的 `k` 位二进制字符。可以用 `map` 存储不同的二进制转换成的十进制数，最后判断 `len(map)` 是否等于 `k` 即可。但是用 `map` 存比较慢，此处换成 `bool` 数组。先构造一个长度为 `k` 的数组，然后每次通过 `mask` 更新这个 `bool` 数组对应十进制的 `bool` 值，并且记录剩余还缺几个二进制数。等剩余的等于 0 的时候，说明所有二进制字符串都出现了，直接输出 `true`，否则循环完以后输出 `false`。

## 代码

```go
package leetcode

import "math"

func hasAllCodes(s string, k int) bool {
	need := int(math.Pow(2.0, float64(k)))
	visited, mask, curr := make([]bool, need), (1<<k)-1, 0
	for i := 0; i < len(s); i++ {
		curr = ((curr << 1) | int(s[i]-'0')) & mask
		if i >= k-1 { // mask 有效位达到了 k 位
			if !visited[curr] {
				need--
				visited[curr] = true
				if need == 0 {
					return true
				}
			}
		}
	}
	return false
}
```