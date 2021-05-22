# [890. Find and Replace Pattern](https://leetcode.com/problems/find-and-replace-pattern/)


## 题目

Given a list of strings `words` and a string `pattern`, return *a list of* `words[i]` *that match* `pattern`. You may return the answer in **any order**.

A word matches the pattern if there exists a permutation of letters `p` so that after replacing every letter `x` in the pattern with `p(x)`, we get the desired word.

Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.

**Example 1:**

```
Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
Output: ["mee","aqq"]
Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation, since a and b map to the same letter.
```

**Example 2:**

```
Input: words = ["a","b","c"], pattern = "a"
Output: ["a","b","c"]
```

**Constraints:**

- `1 <= pattern.length <= 20`
- `1 <= words.length <= 50`
- `words[i].length == pattern.length`
- `pattern` and `words[i]` are lowercase English letters.

## 题目大意

你有一个单词列表 words 和一个模式  pattern，你想知道 words 中的哪些单词与模式匹配。如果存在字母的排列 p ，使得将模式中的每个字母 x 替换为 p(x) 之后，我们就得到了所需的单词，那么单词与模式是匹配的。（回想一下，字母的排列是从字母到字母的双射：每个字母映射到另一个字母，没有两个字母映射到同一个字母。）返回 words 中与给定模式匹配的单词列表。你可以按任何顺序返回答案。

## 解题思路

- 按照题目要求，分别映射两个字符串，words 字符串数组中的字符串与 pattern 字符串每个字母做映射。这里用 map 存储。题目还要求不存在 2 个字母映射到同一个字母的情况，所以再增加一个 map，用来判断当前字母是否已经被映射过了。以上 2 个条件都满足即代表模式匹配上了。最终将所有满足模式匹配的字符串输出即可。

## 代码

```go
package leetcode

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for _, word := range words {
		if match(word, pattern) {
			res = append(res, word)
		}
	}
	return res
}

func match(w, p string) bool {
	if len(w) != len(p) {
		return false
	}
	m, used := make(map[uint8]uint8), make(map[uint8]bool)
	for i := 0; i < len(w); i++ {
		if v, ok := m[p[i]]; ok {
			if w[i] != v {
				return false
			}
		} else {
			if used[w[i]] {
				return false
			}
			m[p[i]] = w[i]
			used[w[i]] = true
		}
	}
	return true
}
```