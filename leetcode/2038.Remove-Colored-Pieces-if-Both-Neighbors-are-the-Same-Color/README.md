# [2038. Remove Colored Pieces if Both Neighbors are the Same Color](https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/)

## 题目

There are n pieces arranged in a line, and each piece is colored either by 'A' or by 'B'. You are given a string colors of length n where colors[i] is the color of the ith piece.

Alice and Bob are playing a game where they take alternating turns removing pieces from the line. In this game, Alice moves first.

- Alice is only allowed to remove a piece colored 'A' if both its neighbors are also colored 'A'. She is not allowed to remove pieces that are colored 'B'.
- Bob is only allowed to remove a piece colored 'B' if both its neighbors are also colored 'B'. He is not allowed to remove pieces that are colored 'A'.
- Alice and Bob cannot remove pieces from the edge of the line.
- If a player cannot make a move on their turn, that player loses and the other player wins.

Assuming Alice and Bob play optimally, return true if Alice wins, or return false if Bob wins.

**Example 1:**

    Input: colors = "AAABABB"
    Output: true
    Explanation:
    AAABABB -> AABABB
    Alice moves first.
    She removes the second 'A' from the left since that is the only 'A' whose neighbors are both 'A'.

    Now it's Bob's turn.
    Bob cannot make a move on his turn since there are no 'B's whose neighbors are both 'B'.
    Thus, Alice wins, so return true.

**Example 2:**

    Input: colors = "AA"
    Output: false
    Explanation:
    Alice has her turn first.
    There are only two 'A's and both are on the edge of the line, so she cannot move on her turn.
    Thus, Bob wins, so return false.

**Example 3:**

    Input: colors = "ABBBBBBBAAA"
    Output: false
    Explanation:
    ABBBBBBBAAA -> ABBBBBBBAA
    Alice moves first.
    Her only option is to remove the second to last 'A' from the right.

    ABBBBBBBAA -> ABBBBBBAA
    Next is Bob's turn.
    He has many options for which 'B' piece to remove. He can pick any.

    On Alice's second turn, she has no more pieces that she can remove.
    Thus, Bob wins, so return false.

**Constraints:**

- 1 <= colors.length <= 100000
- colors consists of only the letters 'A' and 'B'

## 题目大意

总共有 n 个颜色片段排成一列，每个颜色片段要么是 'A' 要么是 'B' 。给你一个长度为 n 的字符串 colors ，其中 colors[i] 表示第 i 个颜色片段的颜色。

Alice 和 Bob 在玩一个游戏，他们轮流从这个字符串中删除颜色。Alice 先手。

- 如果一个颜色片段为 'A' 且相邻两个颜色都是颜色 'A'，那么 Alice 可以删除该颜色片段。Alice不可以删除任何颜色 'B' 片段。
- 如果一个颜色片段为 'B'且相邻两个颜色都是颜色 'B' ，那么 Bob 可以删除该颜色片段。Bob 不可以删除任何颜色 'A' 片段。
- Alice 和 Bob 不能从字符串两端删除颜色片段。
- 如果其中一人无法继续操作，则该玩家 输掉游戏且另一玩家 获胜。

假设 Alice 和 Bob 都采用最优策略，如果 Alice 获胜，请返回true，否则 Bob 获胜，返回false。

## 解题思路

- 统计 Alice 和 Bob 分别可以操作的次数记为 As，Bs 
- 因为 Alice 先手，所以只要 As 大于 Bs，Alice 获胜返回 true，否则 Bob 获胜返回 false

# 代码

```go
package leetcode

func winnerOfGame(colors string) bool {
	As, Bs := 0, 0
	Acont, Bcont := 0, 0
	for _, color := range colors {
		if color == 'A' {
			Acont += 1
			Bcont = 0
		} else {
			Bcont += 1
			Acont = 0
		}
		if Acont >= 3 {
			As += Acont - 2
		}
		if Bcont >= 3 {
			Bs += Bcont - 2
		}
	}
	if As > Bs {
		return true
	}
	return false
}
```