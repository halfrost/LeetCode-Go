# [1145. Binary Tree Coloring Game](https://leetcode.com/problems/binary-tree-coloring-game/)



## 题目

Two players play a turn based game on a binary tree. We are given the `root` of this binary tree, and the number of nodes `n` in the tree. `n` is odd, and each node has a distinct value from `1` to `n`.

Initially, the first player names a value `x` with `1 <= x <= n`, and the second player names a value `y` with `1 <= y <= n` and `y != x`. The first player colors the node with value `x` red, and the second player colors the node with value `y` blue.

Then, the players take turns starting with the first player. In each turn, that player chooses a node of their color (red if player 1, blue if player 2) and colors an **uncolored** neighbor of the chosen node (either the left child, right child, or parent of the chosen node.)

If (and only if) a player cannot choose such a node in this way, they must pass their turn. If both players pass their turn, the game ends, and the winner is the player that colored more nodes.

You are the second player. If it is possible to choose such a `y` to ensure you win the game, return `true`. If it is not possible, return `false`.

**Example 1**:

![https://assets.leetcode.com/uploads/2019/08/01/1480-binary-tree-coloring-game.png](https://assets.leetcode.com/uploads/2019/08/01/1480-binary-tree-coloring-game.png)

```
Input: root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
Output: true
Explanation: The second player can choose the node with value 2.
```

**Constraints**:

- `root` is the root of a binary tree with `n` nodes and distinct node values from `1` to `n`.
- `n` is odd.
- `1 <= x <= n <= 100`

## 题目大意

有两位极客玩家参与了一场「二叉树着色」的游戏。游戏中，给出二叉树的根节点 root，树上总共有 n 个节点，且 n 为奇数，其中每个节点上的值从 1 到 n 各不相同。游戏从「一号」玩家开始（「一号」玩家为红色，「二号」玩家为蓝色），最开始时，

- 「一号」玩家从 [1, n] 中取一个值 x（1 <= x <= n）；
- 「二号」玩家也从 [1, n] 中取一个值 y（1 <= y <= n）且 y != x。
- 「一号」玩家给值为 x 的节点染上红色，而「二号」玩家给值为 y 的节点染上蓝色。

之后两位玩家轮流进行操作，每一回合，玩家选择一个他之前涂好颜色的节点，将所选节点一个 未着色 的邻节点（即左右子节点、或父节点）进行染色。如果当前玩家无法找到这样的节点来染色时，他的回合就会被跳过。若两个玩家都没有可以染色的节点时，游戏结束。着色节点最多的那位玩家获得胜利 ✌️。现在，假设你是「二号」玩家，根据所给出的输入，假如存在一个 y 值可以确保你赢得这场游戏，则返回 true；若无法获胜，就请返回 false。


提示：

- 二叉树的根节点为 root，树上由 n 个节点，节点上的值从 1 到 n 各不相同。
- n 为奇数。
- 1 <= x <= n <= 100

## 解题思路

- 2 个人参加二叉树着色游戏。二叉树节点数为奇数。1 号玩家和 2 号玩家分别在二叉树上选项一个点着色。每一回合，玩家选择一个他之前涂好颜色的节点，将所选节点一个 未着色 的邻节点（即左右子节点、或父节点）进行染色。当有人不能选点着色的时候，他的那个回合会被跳过。双方都没法继续着色的时候游戏结束。着色多的人获胜。问二号玩家是否存在必胜策略？

![](https://img.halfrost.com/Leetcode/leetcode_1145.png)

- 如图所示，当一号玩家选择了一个红色的结点，可能会将二叉树切割为 3 个部分（连通分量），如果选择的是根结点，则可能是 2 个部分或 1 个部分，如果选择叶结点，则是 1 个部分。不过无论哪种情况都无关紧要，我们都可以当成 3 个部分来对待，例如一号玩家选择了一个叶结点，我们也可以把叶结点的左右两个空指针看成大小为 0 的两个部分。
- 那么二号玩家怎样选择蓝色结点才是最优呢？答案是：选择离红色结点最近，且所属连通分量规模最大的那个点。也就是示例图中的 1 号结点。如果我们选择了 1 号结点为蓝色结点，那么可以染成红色的点就只剩下 6 号点和 7 号点了，而蓝色可以把根结点和其左子树全部占据。
- 如何确定蓝色是否有必胜策略，就可以转换为，被红色点切割的三个连通分量中，是否存在一个连通分量，大小大于所有结点数目的一半。统计三个连通分量大小的过程，可以用深度优先搜索（DFS）来实现。当遍历到某一结点，其结点值等于选定的红色结点时，我们统计这个结点的左子树 `red_left` 和右子树 `red_right` 的大小，那么我们就已经找到两个连通分量的大小了，最后一个父结点连通分量的大小，可以用结点总数减去这两个连通分量大小，再减去红色所占结点，即 `parent = n - red_left - red_right - 1`。

## 代码

```go

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var left, right int
	dfsBtreeGameWinningMove(root, &left, &right, x)
	up := n - left - right - 1
	n /= 2
	return left > n || right > n || up > n
}

func dfsBtreeGameWinningMove(node *TreeNode, left, right *int, x int) int {
	if node == nil {
		return 0
	}
	l, r := dfsBtreeGameWinningMove(node.Left, left, right, x), dfsBtreeGameWinningMove(node.Right, left, right, x)
	if node.Val == x {
		*left, *right = l, r
	}
	return l + r + 1
}
```