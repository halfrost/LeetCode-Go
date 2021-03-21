# [1704. Determine if String Halves Are Alike](https://leetcode.com/problems/determine-if-string-halves-are-alike/)

## 题目

You are given a string `s` of even length. Split this string into two halves of equal lengths, and let `a` be the first half and `b` be the second half.

Two strings are **alike** if they have the same number of vowels (`'a'`, `'e'`, `'i'`, `'o'`, `'u'`, `'A'`, `'E'`, `'I'`, `'O'`, `'U'`). Notice that `s` contains uppercase and lowercase letters.

Return `true` *if* `a` *and* `b` *are **alike***. Otherwise, return `false`.

**Example 1:**

```
Input: s = "book"
Output: true
Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel. Therefore, they are alike.
```

**Example 2:**

```
Input: s = "textbook"
Output: false
Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2. Therefore, they are not alike.
Notice that the vowel o is counted twice.
```

**Example 3:**

```
Input: s = "MerryChristmas"
Output: false
```

**Example 4:**

```
Input: s = "AbCdEfGh"
Output: true
```

**Constraints:**

- `2 <= s.length <= 1000`
- `s.length` is even.
- `s` consists of **uppercase and lowercase** letters.

## 题目大意

给你一个偶数长度的字符串 s 。将其拆分成长度相同的两半，前一半为 a ，后一半为 b 。两个字符串 相似 的前提是它们都含有相同数目的元音（'a'，'e'，'i'，'o'，'u'，'A'，'E'，'I'，'O'，'U'）。注意，s 可能同时含有大写和小写字母。如果 a 和 b 相似，返回 true ；否则，返回 false 。

## 解题思路

- 简单题。依题意，分别统计前半段元音字母的个数和后半段元音字母的个数，个数相同则输出 true，不同就输出 false。

## 代码

```go
package leetcode

func halvesAreAlike(s string) bool {
	return numVowels(s[len(s)/2:]) == numVowels(s[:len(s)/2])
}

func numVowels(x string) int {
	res := 0
	for _, c := range x {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			res++
		}
	}
	return res
}
```