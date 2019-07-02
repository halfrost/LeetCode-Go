# [1079. Letter Tile Possibilities](https://leetcode.com/problems/letter-tile-possibilities/)


## 题目

You have a set of `tiles`, where each tile has one letter `tiles[i]` printed on it. Return the number of possible non-empty sequences of letters you can make.

**Example 1:**

    Input: "AAB"
    Output: 8
    Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".

**Example 2:**

    Input: "AAABBC"
    Output: 188

**Note:**

1. `1 <= tiles.length <= 7`
2. `tiles` consists of uppercase English letters.

## 题目大意

你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。提示：  

- 1 <= tiles.length <= 7
- tiles 由大写英文字母组成

## 解题思路

- 题目要求输出所有非空字母序列的数目。这一题是排列和组合的结合题目。组合是可以选择一个字母，二个字母，…… n 个字母。每个组合内是排列问题。比如选择 2 个字母，字母之间相互排序不同是影响最终结果的，不同的排列顺序是不同的解。
- 这道题目由于不需要输出所有解，所以解法可以优化，例如我们在递归计算解的时候，不需要真的遍历原字符串，只需要累加一些字母的频次就可以。当然如果要输出所有解，就需要真实遍历原字符串了(见解法二)。简单的做法是每次递归按照频次累加。因为每次增加一个字母一定是 26 个大写字母中的一个。这里需要注意的是，增加的只能是 26 个字母里面还能取出“机会”的字母，例如递归到到第 3 轮了，A 用完了，这个时候只能取频次还不为 0 的字母拼上去。

![](https://pic.cdn.lizenghai.com/uploads/2019/06/15604249050330956_414598-48d6698970379275.png?x-oss-process=style%2Ffull)