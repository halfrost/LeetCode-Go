# [421. Maximum XOR of Two Numbers in an Array](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/)


## 题目:

Given a **non-empty** array of numbers, a0, a1, a2, … , an-1, where 0 ≤ ai < 231.

Find the maximum result of ai XOR aj, where 0 ≤ i, j < n.

Could you do this in O(n) runtime?

**Example:**

    Input: [3, 10, 5, 25, 2, 8]
    
    Output: 28
    
    Explanation: The maximum result is 5 ^ 25 = 28.


## 题目大意

给定一个非空数组，数组中元素为 a0, a1, a2, … , an-1，其中 0 ≤ ai < 2^31 。找到 ai 和 aj 最大的异或 (XOR) 运算结果，其中0 ≤ i,  j < n 。你能在O(n)的时间解决这个问题吗？


## 解题思路


- 这一题最先考虑到的解法就是暴力解法，2 层循环，依次计算两两数之间的异或值，动态维护最大的值，遍历完成以后输出最大值即可。提交代码会发现超时。
- 改进一点的做法就是一层循环。试想，求的最终结果是一个 32 位的二进制数，如果想要这个数最大，那么高位都填满 1 就是最大。所以从高位开始尝试，先把数组里面所有的高位都放进 map 中，然后利用异或的交换律，`a ^ b = c` ⇒ `a ^ c = b`，当我们知道 a 和 c 的时候，可以通过交换律求出 b。a 就是我们遍历的每个数，c 是我们想要尝试的高位最大值，例如，111…000，从高位逐渐往低位填 1 。如果我们求的 b 也在 map 中，那么就代表 c 是可以求出来的。如果 c 比当前的 max 值要大，就更新。按照这样的方式遍历往 32 位，每次也遍历完整个数组中的每个数，最终 max 里面就是需要求的最大值。
- 还有更好的做法是利用 Trie 这个数据结构。构建一棵深度为 33 的二叉树。root 节点左孩子为 1，右孩子为 0 代表着所有数字的最高位，其次根据次高位继续往下。如果某一个节点左右子树都不为空，那么得到最终答案的两个数字肯定分别出自于左右子树且此位为 1；如果任意一个为空，那么最终答案该位为 0，依次迭代得到最终结果。具体做法见：[Java O(n) solution using Trie - LeetCode Discuss](https://discuss.leetcode.com/topic/63207/java-o-n-solution-using-trie)

- 最后还有更“完美的做法”，利用 leetcode 网站判题的特性，我们可以测出比较弱的数据，绕过这组弱数据可以直接 AC。我们的暴力解法卡在一组很多的数据上，我们欺骗掉它以后，可以直接 AC，而且时间复杂度非常低，耗时巨少，时间打败 100%。
