# [1664. Ways to Make a Fair Array](https://leetcode.com/problems/ways-to-make-a-fair-array/)


## 题目

You are given an integer array `nums`. You can choose **exactly one** index (**0-indexed**) and remove the element. Notice that the index of the elements may change after the removal.

For example, if `nums = [6,1,7,4,1]`:

- Choosing to remove index `1` results in `nums = [6,7,4,1]`.
- Choosing to remove index `2` results in `nums = [6,1,4,1]`.
- Choosing to remove index `4` results in `nums = [6,1,7,4]`.

An array is **fair** if the sum of the odd-indexed values equals the sum of the even-indexed values.

Return the ***number** of indices that you could choose such that after the removal,* `nums` *is **fair**.*

**Example 1:**

```
Input: nums = [2,1,6,4]
Output: 1
Explanation:
Remove index 0: [1,6,4] -> Even sum: 1 + 4 = 5. Odd sum: 6. Not fair.
Remove index 1: [2,6,4] -> Even sum: 2 + 4 = 6. Odd sum: 6. Fair.
Remove index 2: [2,1,4] -> Even sum: 2 + 4 = 6. Odd sum: 1. Not fair.
Remove index 3: [2,1,6] -> Even sum: 2 + 6 = 8. Odd sum: 1. Not fair.
There is 1 index that you can remove to make nums fair.
```

**Example 2:**

```
Input: nums = [1,1,1]
Output: 3
Explanation: You can remove any index and the remaining array is fair.
```

**Example 3:**

```
Input: nums = [1,2,3]
Output: 0
Explanation: You cannot make a fair array after removing any index.
```

**Constraints:**

- `1 <= nums.length <= 105`
- `1 <= nums[i] <= 104`

## 题目大意

给你一个整数数组 nums 。你需要选择 恰好 一个下标（下标从 0 开始）并删除对应的元素。请注意剩下元素的下标可能会因为删除操作而发生改变。

比方说，如果 nums = [6,1,7,4,1] ，那么：

- 选择删除下标 1 ，剩下的数组为 nums = [6,7,4,1] 。
- 选择删除下标 2 ，剩下的数组为 nums = [6,1,4,1] 。
- 选择删除下标 4 ，剩下的数组为 nums = [6,1,7,4] 。

如果一个数组满足奇数下标元素的和与偶数下标元素的和相等，该数组就是一个 平衡数组 。请你返回删除操作后，剩下的数组 nums 是 平衡数组 的 方案数 。

## 解题思路

- 给定一个数组 nums，要求输出仅删除一个元素以后能使得整个数组平衡的方案数。平衡的定义是奇数下标元素总和等于偶数下标元素总和。
- 这一题如果暴力解答，会超时。原因是每次删除元素以后，都重新计算奇偶数位总和比较耗时。应该利用前面计算过的累加和，推导出此次删除元素以后的情况。这样修改以后就不超时了。具体的，如果删除的是元素是奇数位，这个下标的前缀和不变，要变化的是后面的。删除元素后面，原来偶数位的总和变成了奇数位了，原来奇数位的总和变成偶数位了。删除元素后面这半段的总和可以用前缀和计算出来，奇数位的总和减去删除元素的前缀和，就得到了删除元素后面的后缀和。通过这个办法就可以得到删除元素后面的，奇数位总和，偶数位总和。注意这个后缀和是包含了删除元素的。所以最后需要判断删除元素是奇数位还是偶数位，如果是奇数位，那么在计算出来的偶数和上再减去这个删除元素；如果是偶数位，就在计算出来的奇数和上再减去这个删除元素。代码见解法二。
- 这一题还有一种更简洁的写法，就是解法一了。通过了解法二的思考，我们可以知道，每次变换以后的操作可以抽象出来，即三步，减去一个数，判断是否相等，再加上一个数。只不过这三步在解法二中都去判断了奇偶性。如果我们不判断奇偶性，那么代码就可以写成解法一的样子。为什么可以不用管奇偶性呢？因为每次删除一个元素以后，下次再删除，奇偶就发生颠倒了，上次的奇数和到了下次就是偶数和了。想通这一点就可以把代码写成解法一的样子。

## 代码

```go
// 解法一 超简洁写法
func waysToMakeFair(nums []int) int {
	sum, res := [2]int{}, 0
	for i := 0; i < len(nums); i++ {
		sum[i%2] += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		sum[i%2] -= nums[i]
		if sum[i%2] == sum[1-(i%2)] {
			res++
		}
		sum[1-(i%2)] += nums[i]
	}
	return res
}

// 解法二 前缀和，后缀和
func waysToMakeFair1(nums []int) int {
	evenPrefix, oddPrefix, evenSuffix, oddSuffix, res := 0, 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenSuffix += nums[i]
		} else {
			oddSuffix += nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenSuffix -= nums[i]
		} else {
			oddSuffix -= nums[i]
		}
		if (evenPrefix + oddSuffix) == (oddPrefix + evenSuffix) {
			res++
		}
		if i%2 == 0 {
			evenPrefix += nums[i]
		} else {
			oddPrefix += nums[i]
		}
	}
	return res
}
```