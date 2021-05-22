# [810. Chalkboard XOR Game](https://leetcode.com/problems/chalkboard-xor-game/)


## 题目

We are given non-negative integers nums[i] which are written on a chalkboard.  Alice and Bob take turns erasing exactly one number from the chalkboard, with Alice starting first.  If erasing a number causes the bitwise XOR of all the elements of the chalkboard to become 0, then that player loses.  (Also, we'll say the bitwise XOR of one element is that element itself, and the bitwise XOR of no elements is 0.)

Also, if any player starts their turn with the bitwise XOR of all the elements of the chalkboard equal to 0, then that player wins.

Return True if and only if Alice wins the game, assuming both players play optimally.

```
Example:Input: nums = [1, 1, 2]
Output: false
Explanation:
Alice has two choices: erase 1 or erase 2.
If she erases 1, the nums array becomes [1, 2]. The bitwise XOR of all the elements of the chalkboard is 1 XOR 2 = 3. Now Bob can remove any element he wants, because Alice will be the one to erase the last element and she will lose.
If Alice erases 2 first, now nums becomes [1, 1]. The bitwise XOR of all the elements of the chalkboard is 1 XOR 1 = 0. Alice will lose.
```

**Notes:**

- `1 <= N <= 1000`.
- `0 <= nums[i] <= 2^16`.

## 题目大意

黑板上写着一个非负整数数组 nums[i] 。Alice 和 Bob 轮流从黑板上擦掉一个数字，Alice 先手。如果擦除一个数字后，剩余的所有数字按位异或运算得出的结果等于 0 的话，当前玩家游戏失败。 (另外，如果只剩一个数字，按位异或运算得到它本身；如果无数字剩余，按位异或运算结果为 0。）并且，轮到某个玩家时，如果当前黑板上所有数字按位异或运算结果等于 0，这个玩家获胜。假设两个玩家每步都使用最优解，当且仅当 Alice 获胜时返回 true。

## 解题思路

- Alice 必胜情况之一，Alice 先手，起始数组全部元素本身异或结果就为 0 。不需要擦除数字便自动获胜。除去这个情况，还有其他情况么？由于 2 人是交替擦除数字，且每次都恰好擦掉一个数字，因此对于这两人中的任意一人，其每次在擦除数字前，黑板上剩余数字的个数的奇偶性一定都是相同的。于是奇偶性成为突破口。
- 如果 nums 的长度是偶数，Alice 先手是否必败呢？如果必败，代表无论擦掉哪一个数字，剩余所有数字的异或结果都等于 0。利用反证法证明上述结论是错误的。首先 $num[0] \oplus num[1] \oplus num[2] \oplus \cdots \oplus num[n-1] = X ≠ 0$，初始所有元素异或结果不为 0。假设 Alice 当前擦掉第 i 个元素，0 ≤ i < n。令 $X_{n}$ 代表擦掉第 n 位元素以后剩余元素异或的结果。由证题，无论擦掉哪一个数字，剩余所有数字的异或结果都等于 0。所以 $X_{0} \oplus X_{1} \oplus X_{2} \oplus \cdots  \oplus X_{n-1} = 0$。

    $$\begin{aligned}0 &= X_{0} \oplus  X_{1} \oplus X_{2} \oplus \cdots  \oplus X_{n-1} \\0 &= (X \oplus nums[0]) \oplus (X \oplus nums[1]) \oplus (X \oplus nums[2]) \oplus \cdots  \oplus (X \oplus nums[n-1])\\ 0 &= (X \oplus X \oplus \cdots  \oplus X) \oplus (nums[0] \oplus nums[1] \oplus nums[2] \oplus \cdots  \oplus nums[n-1])\\0 &= 0 \oplus X\\\\\Rightarrow X &= 0\\\end{aligned}$$

    由于 n 为偶数，所以 n 个 X 的异或结果为 0。最终推出 X = 0，很明显与前提 X ≠ 0 冲突。所以原命题，代表无论擦掉哪一个数字，剩余所有数字的异或结果都等于 0 是错误的。也就是说，当 n 为偶数时，代表无论擦掉哪一个数字，剩余所有数字的异或结果都不等于 0。即 Alice 有必胜策略。换句话说，当数组的长度是偶数时，先手 Alice 总能找到一个数字，在擦掉这个数字之后剩余的所有数字异或结果不等于 0。

- 综上，Alice 必胜策略有 2 种情况：
    1. 数组 nums 的全部元素初始本身异或结果就等于 0。
    2. 数组 nums 的长度是偶数。

## 代码

```go
package leetcode

func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0
}
```