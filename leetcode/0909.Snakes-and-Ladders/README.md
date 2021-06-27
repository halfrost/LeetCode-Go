# [909. Snakes and Ladders](https://leetcode.com/problems/snakes-and-ladders/)


## 题目

On an N x N `board`, the numbers from `1` to `N*N` are written *boustrophedonically* **starting from the bottom left of the board**, and alternating direction each row.  For example, for a 6 x 6 board, the numbers are written as follows:


![](https://assets.leetcode.com/uploads/2018/09/23/snakes.png)

You start on square `1` of the board (which is always in the last row and first column).  Each move, starting from square `x`, consists of the following:

- You choose a destination square `S` with number `x+1`, `x+2`, `x+3`, `x+4`, `x+5`, or `x+6`, provided this number is `<= N*N`.
    - (This choice simulates the result of a standard 6-sided die roll: ie., there are always **at most 6 destinations, regardless of the size of the board**.)
- If `S` has a snake or ladder, you move to the destination of that snake or ladder. Otherwise, you move to `S`.

A board square on row `r` and column `c` has a "snake or ladder" if `board[r][c] != -1`.  The destination of that snake or ladder is `board[r][c]`.

Note that you only take a snake or ladder at most once per move: if the destination to a snake or ladder is the start of another snake or ladder, you do **not** continue moving.  (For example, if the board is `[[4,-1],[-1,3]]`, and on the first move your destination square is `2`, then you finish your first move at `3`, because you do **not** continue moving to `4`.)

Return the least number of moves required to reach square N*N.  If it is not possible, return `-1`.

**Example 1:**

```
Input:[
[-1,-1,-1,-1,-1,-1],
[-1,-1,-1,-1,-1,-1],
[-1,-1,-1,-1,-1,-1],
[-1,35,-1,-1,13,-1],
[-1,-1,-1,-1,-1,-1],
[-1,15,-1,-1,-1,-1]]
Output:4
Explanation:
At the beginning, you start at square 1 [at row 5, column 0].
You decide to move to square 2, and must take the ladder to square 15.
You then decide to move to square 17 (row 3, column 5), and must take the snake to square 13.
You then decide to move to square 14, and must take the ladder to square 35.
You then decide to move to square 36, ending the game.
It can be shown that you need at least 4 moves to reach the N*N-th square, so the answer is 4.

```

**Note:**

1. `2 <= board.length = board[0].length <= 20`
2. `board[i][j]` is between `1` and `N*N` or is equal to `1`.
3. The board square with number `1` has no snake or ladder.
4. The board square with number `N*N` has no snake or ladder.

## 题目大意

N x N 的棋盘 board 上，按从 1 到 N*N 的数字给方格编号，编号 从左下角开始，每一行交替方向。r 行 c 列的棋盘，按前述方法编号，棋盘格中可能存在 “蛇” 或 “梯子”；如果 board[r][c] != -1，那个蛇或梯子的目的地将会是 board[r][c]。玩家从棋盘上的方格 1 （总是在最后一行、第一列）开始出发。每一回合，玩家需要从当前方格 x 开始出发，按下述要求前进：选定目标方格：

- 选择从编号 x+1，x+2，x+3，x+4，x+5，或者 x+6 的方格中选出一个目标方格 s ，目标方格的编号 <= N*N。该选择模拟了掷骰子的情景，无论棋盘大小如何，你的目的地范围也只能处于区间 [x+1, x+6] 之间。
- 传送玩家：如果目标方格 S 处存在蛇或梯子，那么玩家会传送到蛇或梯子的目的地。否则，玩家传送到目标方格 S。

注意，玩家在每回合的前进过程中最多只能爬过蛇或梯子一次：就算目的地是另一条蛇或梯子的起点，你也不会继续移动。返回达到方格 N*N 所需的最少移动次数，如果不可能，则返回 -1。

## 解题思路

- 这一题可以抽象为在有向图上求下标 1 的起点到下标 `N^2` 的终点的最短路径。用广度优先搜索。棋盘可以抽象成一个包含 `N^2` 个节点的有向图，对于每个节点 `x`，若 `x+i (1 ≤ i ≤ 6)` 上没有蛇或梯子，则连一条从 `x` 到 `x+i` 的有向边；否则记蛇梯的目的地为 `y`，连一条从 `x` 到 `y` 的有向边。然后按照最短路径的求解方式便可解题。时间复杂度 O(n^2)，空间复杂度 O(n^2)。
- 此题棋盘上的下标是蛇形的，所以遍历下一个点的时候需要转换坐标。具体做法根据行的奇偶性，行号为偶数，下标从左往右，行号为奇数，下标从右往左。具体实现见 `getRowCol()` 函数。

## 代码

```go
package leetcode

type pair struct {
	id, step int
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([]bool, n*n+1)
	queue := []pair{{1, 0}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for i := 1; i <= 6; i++ {
			nxt := p.id + i
			if nxt > n*n { // 超出边界
				break
			}
			r, c := getRowCol(nxt, n) // 得到下一步的行列
			if board[r][c] > 0 {      // 存在蛇或梯子
				nxt = board[r][c]
			}
			if nxt == n*n { // 到达终点
				return p.step + 1
			}
			if !visited[nxt] {
				visited[nxt] = true
				queue = append(queue, pair{nxt, p.step + 1}) // 扩展新状态
			}
		}
	}
	return -1
}

func getRowCol(id, n int) (r, c int) {
	r, c = (id-1)/n, (id-1)%n
	if r%2 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return r, c
}
```