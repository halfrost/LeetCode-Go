# [1178. Number of Valid Words for Each Puzzle](https://leetcode.com/problems/number-of-valid-words-for-each-puzzle/)


## 题目

With respect to a given `puzzle` string, a `word` is *valid* if both the following conditions are satisfied:

- `word` contains the first letter of `puzzle`.
- For each letter in `word`, that letter is in `puzzle`.For example, if the puzzle is "abcdefg", then valid words are "faced", "cabbage", and "baggage"; while invalid words are "beefed" (doesn't include "a") and "based" (includes "s" which isn't in the puzzle).

Return an array `answer`, where `answer[i]` is the number of words in the given word list `words` that are valid with respect to the puzzle `puzzles[i]`.

**Example :**

```
Input: 
words = ["aaaa","asas","able","ability","actt","actor","access"], 
puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
Output: [1,1,3,2,4,0]
Explanation:
1 valid word for "aboveyz" : "aaaa" 
1 valid word for "abrodyz" : "aaaa"
3 valid words for "abslute" : "aaaa", "asas", "able"
2 valid words for "absoryz" : "aaaa", "asas"
4 valid words for "actresz" : "aaaa", "asas", "actt", "access"
There're no valid words for "gaswxyz" cause none of the words in the list contains letter 'g'.

```

**Constraints:**

- `1 <= words.length <= 10^5`
- `4 <= words[i].length <= 50`
- `1 <= puzzles.length <= 10^4`
- `puzzles[i].length == 7`
- `words[i][j]`, `puzzles[i][j]` are English lowercase letters.
- Each `puzzles[i]` doesn't contain repeated characters.

## 题目大意

外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。

字谜的迷面 puzzle 按字符串形式给出，如果一个单词 word 符合下面两个条件，那么它就可以算作谜底：

- 单词 word 中包含谜面 puzzle 的第一个字母。
- 单词 word 中的每一个字母都可以在谜面 puzzle 中找到。
例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage", 和 "baggage"；而 "beefed"（不含字母 "a"）以及 "based"（其中的 "s" 没有出现在谜面中）都不能作为谜底。

返回一个答案数组 answer，数组中的每个元素 answer[i] 是在给出的单词列表 words 中可以作为字谜迷面 puzzles[i] 所对应的谜底的单词数目。

提示：

- 1 <= words.length <= 10^5
- 4 <= words[i].length <= 50
- 1 <= puzzles.length <= 10^4
- puzzles[i].length == 7
- words[i][j], puzzles[i][j] 都是小写英文字母。
- 每个 puzzles[i] 所包含的字符都不重复。

## 解题思路

- 首先题目中两个限制条件非常关键：**puzzles[i].length == 7**，**每个 puzzles[i] 所包含的字符都不重复**。也就是说穷举每个puzzle的子串的搜索空间就是2^7=128，而且不用考虑去重问题。
- 因为谜底的判断只跟字符是否出现有关，跟字符的个数无关，另外都是小写的英文字母，所以可以用 `bitmap` 来表示单词(word)。
- 利用 `map` 记录不同状态的单词(word)的个数。
- 根据题意，如果某个单词(word)是某个字谜(puzzle)的谜底，那么 `word` 的 `bitmap` 肯定对应于 `puzzle` 某个子串的 `bitmap` 表示，且 `bitmap` 中包含 `puzzle` 的第一个字母的 `bit` 占用。
- 问题就转换为：求每一个 `puzzle` 的每一个子串，然后求和这个子串具有相同 `bitmap` 表示且 `word` 中包含 `puzzle` 的第一个字母的 `word` 的个数。

## 代码

```go
package leetcode

/*
	匹配跟单词中的字母顺序，字母个数都无关，可以用bitmap压缩
	1. 记录word中 利用map记录各种bit标示的个数
	2. puzzles 中各个字母都不相同! 记录bitmap，然后搜索子空间中各种bit标识的个数的和
	   因为puzzles长度最长是7，所以搜索空间 2^7
*/
func findNumOfValidWords(words []string, puzzles []string) []int {
	wordBitStatusMap, res := make(map[uint32]int, 0), []int{}
	for _, w := range words {
		wordBitStatusMap[toBitMap([]byte(w))]++
	}
	for _, p := range puzzles {
		var bitMap uint32
		var totalNum int
		bitMap |= (1 << (p[0] - 'a')) //work中要包含 p 的第一个字母 所以这个bit位上必须是1
		findNum([]byte(p)[1:], bitMap, &totalNum, wordBitStatusMap)
		res = append(res, totalNum)
	}
	return res
}

func toBitMap(word []byte) uint32 {
	var res uint32
	for _, b := range word {
		res |= (1 << (b - 'a'))
	}
	return res
}

//利用dfs 搜索 pussles的子空间
func findNum(puzzles []byte, bitMap uint32, totalNum *int, m map[uint32]int) {
	if len(puzzles) == 0 {
		*totalNum = *totalNum + m[bitMap]
		return
	}
	//不包含puzzles[0],即puzzles[0]对应bit是0
	findNum(puzzles[1:], bitMap, totalNum, m)
	//包含puzzles[0],即puzzles[0]对应bit是1
	bitMap |= (1 << (puzzles[0] - 'a'))
	findNum(puzzles[1:], bitMap, totalNum, m)
	bitMap ^= (1 << (puzzles[0] - 'a')) //异或 清零
	return
}
```