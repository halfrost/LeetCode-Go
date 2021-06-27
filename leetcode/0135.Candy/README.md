# [135. Candy](https://leetcode.com/problems/candy/)


## 题目

There are `n` children standing in a line. Each child is assigned a rating value given in the integer array `ratings`.

You are giving candies to these children subjected to the following requirements:

- Each child must have at least one candy.
- Children with a higher rating get more candies than their neighbors.

Return *the minimum number of candies you need to have to distribute the candies to the children*.

**Example 1:**

```
Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
```

**Example 2:**

```
Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.
```

**Constraints:**

- `n == ratings.length`
- `1 <= n <= 2 * 10^4`
- `0 <= ratings[i] <= 2 * 10^4`

## 题目大意

老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。你需要按照以下要求，帮助老师给这些孩子分发糖果：

- 每个孩子至少分配到 1 个糖果。
- 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。

那么这样下来，老师至少需要准备多少颗糖果呢？

## 解题思路

- 本题的突破口在于，评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果，这句话。这个规则可以理解为 2 条规则，想象成按身高排队，站在下标为 0 的地方往后“看”，评分高即为个子高的，应该比前面个子矮(评分低)的分到糖果多；站在下标为 n - 1 的地方往后“看”，评分高即为个子高的，同样应该比前面个子矮(评分低)的分到糖果多。你可能会有疑问，规则都是一样的，为什么会出现至少需要多少糖果呢？因为可能出现评分一样高的同学。扫描数组两次，处理出每一个学生分别满足左规则或右规则时，最少需要被分得的糖果数量。每个人最终分得的糖果数量即为这两个数量的最大值。两次遍历结束，将所有糖果累加起来即为至少需要准备的糖果数。由于每个人至少分配到 1 个糖果，所以每个人糖果数再加一。

## 代码

```go
package leetcode

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] += candies[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	total := 0
	for _, candy := range candies {
		total += candy + 1
	}
	return total
}
```