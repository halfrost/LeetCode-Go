# [1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence](https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/)


## 题目

Given a `sentence` that consists of some words separated by a **single space**, and a `searchWord`.

You have to check if `searchWord` is a prefix of any word in `sentence`.

Return *the index of the word* in `sentence` where `searchWord` is a prefix of this word (**1-indexed**).

If `searchWord` is a prefix of more than one word, return the index of the first word **(minimum index)**. If there is no such word return **-1**.

A **prefix** of a string `S` is any leading contiguous substring of `S`.

**Example 1**:

```
Input: sentence = "i love eating burger", searchWord = "burg"
Output: 4
Explanation: "burg" is prefix of "burger" which is the 4th word in the sentence.

```

**Example 2**:

```
Input: sentence = "this problem is an easy problem", searchWord = "pro"
Output: 2
Explanation: "pro" is prefix of "problem" which is the 2nd and the 6th word in the sentence, but we return 2 as it's the minimal index.

```

**Example 3**:

```
Input: sentence = "i am tired", searchWord = "you"
Output: -1
Explanation: "you" is not a prefix of any word in the sentence.

```

**Example 4**:

```
Input: sentence = "i use triple pillow", searchWord = "pill"
Output: 4

```

**Example 5**:

```
Input: sentence = "hello from the other side", searchWord = "they"
Output: -1

```

**Constraints**:

- `1 <= sentence.length <= 100`
- `1 <= searchWord.length <= 10`
- `sentence` consists of lowercase English letters and spaces.
- `searchWord` consists of lowercase English letters.

## 题目大意

给你一个字符串 sentence 作为句子并指定检索词为 searchWord ，其中句子由若干用 单个空格 分隔的单词组成。请你检查检索词 searchWord 是否为句子 sentence 中任意单词的前缀。

- 如果 searchWord 是某一个单词的前缀，则返回句子 sentence 中该单词所对应的下标（下标从 1 开始）。
- 如果 searchWord 是多个单词的前缀，则返回匹配的第一个单词的下标（最小下标）。
- 如果 searchWord 不是任何单词的前缀，则返回 -1 。

字符串 S 的 「前缀」是 S 的任何前导连续子字符串。

## 解题思路

- 给出 2 个字符串，一个是匹配串，另外一个是句子。在句子里面查找带匹配串前缀的单词，并返回第一个匹配单词的下标。
- 简单题。按照题意，扫描一遍句子，一次匹配即可。

## 代码

```go

package leetcode

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, v := range strings.Split(sentence, " ") {
		if strings.HasPrefix(v, searchWord) {
			return i + 1
		}
	}
	return -1
}

```