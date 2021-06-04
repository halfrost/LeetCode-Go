# [97. Interleaving String](https://leetcode.com/problems/interleaving-string/)


## 题目

Given strings `s1`, `s2`, and `s3`, find whether `s3` is formed by an **interleaving** of `s1` and `s2`.

An **interleaving** of two strings `s` and `t` is a configuration where they are divided into **non-empty** substrings such that:

- `s = s1 + s2 + ... + sn`
- `t = t1 + t2 + ... + tm`
- `|n - m| <= 1`
- The **interleaving** is `s1 + t1 + s2 + t2 + s3 + t3 + ...` or `t1 + s1 + t2 + s2 + t3 + s3 + ...`

**Note:** `a + b` is the concatenation of strings `a` and `b`.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/09/02/interleave.jpg](https://assets.leetcode.com/uploads/2020/09/02/interleave.jpg)

```
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true

```

**Example 2:**

```
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false

```

**Example 3:**

```
Input: s1 = "", s2 = "", s3 = ""
Output: true

```

**Constraints:**

- `0 <= s1.length, s2.length <= 100`
- `0 <= s3.length <= 200`
- `s1`, `s2`, and `s3` consist of lowercase English letters.

**Follow up:** Could you solve it using only `O(s2.length)` additional memory space?

## 题目大意

给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

- s = s1 + s2 + ... + sn
- t = t1 + t2 + ... + tm
- |n - m| <= 1
- 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...

提示：a + b 意味着字符串 a 和 b 连接。

## 解题思路

- 深搜或者广搜暴力解题。笔者用深搜实现的。记录 s1 和 s2 串当前比较的位置 p1 和 p2。如果 s3[p1+p2] 的位置上等于 s1[p1] 或者 s2[p2] 代表能匹配上，那么继续往后移动 p1 和 p2 相应的位置。因为是交错字符串，所以判断匹配的位置是 s3[p1+p2] 的位置。如果仅仅这么写，会超时，s1 和 s2 两个字符串重复交叉判断的位置太多了。需要加上记忆化搜索。可以用 visited[i][j] 这样的二维数组来记录是否搜索过了。笔者为了压缩空间，将 i 和 j 编码压缩到一维数组了。i * len(s3) + j 是唯一下标，所以可以用这种方式存储是否搜索过。具体代码见下面的实现。

## 代码

```go
package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	visited := make(map[int]bool)
	return dfs(s1, s2, s3, 0, 0, visited)
}

func dfs(s1, s2, s3 string, p1, p2 int, visited map[int]bool) bool {
	if p1+p2 == len(s3) {
		return true
	}
	if _, ok := visited[(p1*len(s3))+p2]; ok {
		return false
	}
	visited[(p1*len(s3))+p2] = true
	var match1, match2 bool
	if p1 < len(s1) && s3[p1+p2] == s1[p1] {
		match1 = true
	}
	if p2 < len(s2) && s3[p1+p2] == s2[p2] {
		match2 = true
	}
	if match1 && match2 {
		return dfs(s1, s2, s3, p1+1, p2, visited) || dfs(s1, s2, s3, p1, p2+1, visited)
	} else if match1 {
		return dfs(s1, s2, s3, p1+1, p2, visited)
	} else if match2 {
		return dfs(s1, s2, s3, p1, p2+1, visited)
	} else {
		return false
	}
}
```