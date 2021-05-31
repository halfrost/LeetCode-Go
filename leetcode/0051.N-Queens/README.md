# [51. N-Queens](https://leetcode.com/problems/n-queens/)


## 题目

The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

![](https://assets.leetcode.com/uploads/2018/10/12/8-queens.png)

Given an integer *n*, return all distinct solutions to the *n*-queens puzzle.

Each solution contains a distinct board configuration of the *n*-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space respectively.

**Example:**


    Input: 4
    Output: [
     [".Q..",  // Solution 1
      "...Q",
      "Q...",
      "..Q."],
    
     ["..Q.",  // Solution 2
      "Q...",
      "...Q",
      ".Q.."]
    ]
    Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.


## 题目大意

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。


## 解题思路

- 求解 n 皇后问题
- 利用 col 数组记录列信息，col 有 `n` 列。用 dia1，dia2 记录从左下到右上的对角线，从左上到右下的对角线的信息，dia1 和 dia2 分别都有 `2*n-1` 个。
- dia1 对角线的规律是 `i + j 是定值`，例如[0,0]，为 0；[1,0]、[0,1] 为 1；[2,0]、[1,1]、[0,2] 为 2；
- dia2 对角线的规律是 `i - j 是定值`，例如[0,7]，为 -7；[0,6]、[1,7] 为 -6；[0,5]、[1,6]、[2,7] 为 -5；为了使他们从 0 开始，i - j + n - 1 偏移到 0 开始，所以 dia2 的规律是 `i - j + n - 1 为定值`。
- 还有一个位运算的方法，每行只能选一个位置放皇后，那么对每行遍历可能放皇后的位置。如何高效判断哪些点不能放皇后呢？这里的做法毕竟巧妙，把所有之前选过的点按照顺序存下来，然后根据之前选的点到当前行的距离，就可以快速判断是不是会有冲突。举个例子: 假如在 4 皇后问题中，如果第一二行已经选择了位置 [1, 3]，那么在第三行选择时，首先不能再选 1, 3 列了，而对于第三行， 1 距离长度为2，所以它会影响到 -1, 3 两个列。同理，3 在第二行，距离第三行为 1，所以 3 会影响到列 2, 4。由上面的结果，我们知道 -1, 4 超出边界了不用去管，别的不能选的点是 1, 2, 3，所以第三行就只能选 0。在代码实现中，可以在每次遍历前根据之前选择的情况生成一个 occupied 用来记录当前这一行，已经被选了的和由于之前皇后攻击范围所以不能选的位置，然后只选择合法的位置进入到下一层递归。另外就是预处理了一个皇后放不同位置的字符串，这样这些字符串在返回结果的时候是可以在内存中复用的，省一点内存。

