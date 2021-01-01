# [605. Can Place Flowers](https://leetcode.com/problems/can-place-flowers/)

## 题目

You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in **adjacent** plots.

Given an integer array `flowerbed` containing `0`'s and `1`'s, where `0` means empty and `1` means not empty, and an integer `n`, return *if* `n` new flowers can be planted in the `flowerbed` without violating the no-adjacent-flowers rule.

**Example 1:**

```
Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
```

**Example 2:**

```
Input: flowerbed = [1,0,0,0,1], n = 2
Output: false
```

**Constraints:**

- `1 <= flowerbed.length <= 2 * 104`
- `flowerbed[i]` is `0` or `1`.
- There are no two adjacent flowers in `flowerbed`.
- `0 <= n <= flowerbed.length`

## 题目大意

假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 n 。能否在不打破种植规则的情况下种入 n 朵花？能则返回True，不能则返回False。

## 解题思路

- 这一题最容易想到的解法是步长为 2 遍历数组，依次计数 0 的个数。有 2 种特殊情况需要单独判断，第一种情况是首尾连续多个 0，例如，00001 和 10000，第二种情况是 2 个 1 中间存在的 0 不足以种花，例如，1001 和 100001，1001 不能种任何花，100001 只能种一种花。单独判断出这 2 种情况，这一题就可以 AC 了。
- 换个思路，找到可以种花的基本单元是 00，那么上面那 2 种特殊情况都可以统一成一种情况。判断是否当前存在 00 的组合，如果存在 00 的组合，都可以种花。末尾的情况需要单独判断，如果末尾为 0，也可以种花。这个时候不需要再找 00 组合，因为会越界。代码实现如下，思路很简洁明了。

## 代码

```go
package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	lenth := len(flowerbed)
	for i := 0; i < lenth && n > 0; i += 2 {
		if flowerbed[i] == 0 {
			if i+1 == lenth || flowerbed[i+1] == 0 {
				n--
			} else {
				i++
			}
		}
	}
	if n == 0 {
		return true
	}
	return false
}
```