# [914. X of a Kind in a Deck of Cards](https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/)


## 题目

In a deck of cards, each card has an integer written on it.

Return `true` if and only if you can choose `X >= 2` such that it is possible to split the entire deck into 1 or more groups of cards, where:

- Each group has exactly `X` cards.
- All the cards in each group have the same integer.

**Example 1**:

```
Input: deck = [1,2,3,4,4,3,2,1]
Output: true
Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].
```

**Example 2**:

```
Input: deck = [1,1,1,2,2,2,3,3]
Output: false´
Explanation: No possible partition.
```

**Example 3**:

```
Input: deck = [1]
Output: false
Explanation: No possible partition.
```

**Example 4**:

```
Input: deck = [1,1]
Output: true
Explanation: Possible partition [1,1].
```

**Example 5**:

```
Input: deck = [1,1,2,2,2,2]
Output: true
Explanation: Possible partition [1,1],[2,2],[2,2].
```

**Constraints**:

- `1 <= deck.length <= 10^4`
- `0 <= deck[i] < 10^4`

## 题目大意

给定一副牌，每张牌上都写着一个整数。此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：

- 每组都有 X 张牌。
- 组内所有的牌上都写着相同的整数。

仅当你可选的 X >= 2 时返回 true。


## 解题思路

- 给定一副牌，要求选出数字 X，使得每组都有 X 张牌，每组牌的数字都相同。当 X ≥ 2 的时候，输出 true。
- 通过分析题目，我们可以知道，只有当 X 为所有 count 的约数，即所有 count 的最大公约数的约数时，才存在可能的分组。因此我们只要求出所有 count 的最大公约数 g，判断 g 是否大于等于 2 即可，如果大于等于 2，则满足条件，否则不满足。
- 时间复杂度：O(NlogC)，其中 N 是卡牌的个数，C 是数组 deck 中数的范围，在本题中 C 的值为 10000。求两个数最大公约数的复杂度是 O(logC)，需要求最多 N - 1 次。空间复杂度：O(N + C) 或 O(N)。

## 代码

```go

package leetcode

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	m, g := map[int]int{}, -1
	for _, d := range deck {
		m[d]++
	}
	for _, v := range m {
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g >= 2
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

```