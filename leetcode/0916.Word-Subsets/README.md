# [916. Word Subsets](https://leetcode.com/problems/word-subsets/)


## 题目

We are given two arrays `A` and `B` of words.  Each word is a string of lowercase letters.

Now, say that word `b` is a subset of word `a` ****if every letter in `b` occurs in `a`, **including multiplicity**.  For example, `"wrr"` is a subset of `"warrior"`, but is not a subset of `"world"`.

Now say a word `a` from `A` is *universal* if for every `b` in `B`, `b` is a subset of `a`.

Return a list of all universal words in `A`.  You can return the words in any order.

**Example 1:**

```
Input:A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
Output:["facebook","google","leetcode"]
```

**Example 2:**

```
Input:A = ["amazon","apple","facebook","google","leetcode"], B = ["l","e"]
Output:["apple","google","leetcode"]
```

**Example 3:**

```
Input:A = ["amazon","apple","facebook","google","leetcode"], B = ["e","oo"]
Output:["facebook","google"]
```

**Example 4:**

```
Input:A = ["amazon","apple","facebook","google","leetcode"], B = ["lo","eo"]
Output:["google","leetcode"]
```

**Example 5:**

```
Input:A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
Output:["facebook","leetcode"]
```

**Note:**

1. `1 <= A.length, B.length <= 10000`
2. `1 <= A[i].length, B[i].length <= 10`
3. `A[i]` and `B[i]` consist only of lowercase letters.
4. All words in `A[i]` are unique: there isn't `i != j` with `A[i] == A[j]`.

## 题目大意

我们给出两个单词数组 A 和 B。每个单词都是一串小写字母。现在，如果 b 中的每个字母都出现在 a 中，包括重复出现的字母，那么称单词 b 是单词 a 的子集。 例如，“wrr” 是 “warrior” 的子集，但不是 “world” 的子集。如果对 B 中的每一个单词 b，b 都是 a 的子集，那么我们称 A 中的单词 a 是通用的。你可以按任意顺序以列表形式返回 A 中所有的通用单词。

## 解题思路

- 简单题。先统计出 B 数组中单词每个字母的频次，再在 A 数组中依次判断每个单词是否超过了这个频次，如果超过了即输出。

## 代码

```go
package leetcode

func wordSubsets(A []string, B []string) []string {
	var counter [26]int
	for _, b := range B {
		var m [26]int
		for _, c := range b {
			j := c - 'a'
			m[j]++
		}
		for i := 0; i < 26; i++ {
			if m[i] > counter[i] {
				counter[i] = m[i]
			}
		}
	}
	var res []string
	for _, a := range A {
		var m [26]int
		for _, c := range a {
			j := c - 'a'
			m[j]++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if m[i] < counter[i] {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, a)
		}
	}
	return res
}
```