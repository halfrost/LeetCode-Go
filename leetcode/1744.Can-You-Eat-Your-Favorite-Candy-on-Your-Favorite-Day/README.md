# [1744. Can You Eat Your Favorite Candy on Your Favorite Day?](https://leetcode.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/)

## 题目

You are given a **(0-indexed)** array of positive integers `candiesCount` where `candiesCount[i]` represents the number of candies of the `ith` type you have. You are also given a 2D array `queries` where `queries[i] = [favoriteTypei, favoriteDayi, dailyCapi]`.

You play a game with the following rules:

- You start eating candies on day **`0`**.
- You **cannot** eat **any** candy of type `i` unless you have eaten **all** candies of type `i - 1`.
- You must eat **at least** **one** candy per day until you have eaten all the candies.

Construct a boolean array `answer` such that `answer.length == queries.length` and `answer[i]` is `true` if you can eat a candy of type `favoriteTypei` on day `favoriteDayi` without eating **more than** `dailyCapi` candies on **any** day, and `false` otherwise. Note that you can eat different types of candy on the same day, provided that you follow rule 2.

Return *the constructed array* `answer`.

**Example 1:**

```
Input: candiesCount = [7,4,5,3,8], queries = [[0,2,2],[4,2,4],[2,13,1000000000]]
Output: [true,false,true]
Explanation:
1- If you eat 2 candies (type 0) on day 0 and 2 candies (type 0) on day 1, you will eat a candy of type 0 on day 2.
2- You can eat at most 4 candies each day.
   If you eat 4 candies every day, you will eat 4 candies (type 0) on day 0 and 4 candies (type 0 and type 1) on day 1.
   On day 2, you can only eat 4 candies (type 1 and type 2), so you cannot eat a candy of type 4 on day 2.
3- If you eat 1 candy each day, you will eat a candy of type 2 on day 13.
```

**Example 2:**

```
Input: candiesCount = [5,2,6,4,1], queries = [[3,1,2],[4,10,3],[3,10,100],[4,100,30],[1,3,1]]
Output: [false,true,true,false,false]
```

**Constraints:**

- `1 <= candiesCount.length <= 105`
- `1 <= candiesCount[i] <= 105`
- `1 <= queries.length <= 105`
- `queries[i].length == 3`
- `0 <= favoriteTypei < candiesCount.length`
- `0 <= favoriteDayi <= 109`
- `1 <= dailyCapi <= 109`

## 题目大意

给你一个下标从 0 开始的正整数数组 candiesCount ，其中 candiesCount[i] 表示你拥有的第 i 类糖果的数目。同时给你一个二维数组 queries ，其中 queries[i] = [favoriteTypei, favoriteDayi, dailyCapi] 。你按照如下规则进行一场游戏：

- 你从第 0 天开始吃糖果。
- 你在吃完 所有 第 i - 1 类糖果之前，不能 吃任何一颗第 i 类糖果。
- 在吃完所有糖果之前，你必须每天 至少 吃 一颗 糖果。

请你构建一个布尔型数组 answer ，满足 answer.length == queries.length 。answer[i] 为 true 的条件是：在每天吃 不超过 dailyCapi 颗糖果的前提下，你可以在第 favoriteDayi 天吃到第 favoriteTypei 类糖果；否则 answer[i] 为 false 。注意，只要满足上面 3 条规则中的第二条规则，你就可以在同一天吃不同类型的糖果。请你返回得到的数组 answer 。

## 解题思路

- 每天吃糖个数的下限是 1 颗，上限是 dailyCap。针对每一个 query 查询在第 i 天能否吃到 i 类型的糖果，要想吃到 i 类型的糖果，必须吃完 i-1 类型的糖果。意味着在 [favoriteDayi + 1, (favoriteDayi+1)×dailyCapi] 区间内能否包含一颗第 favoriteTypei 类型的糖果。如果能包含则输出 true，不能包含则输出 false。吃的糖果数是累积的，所以这里利用前缀和计算出累积吃糖果数所在区间 [sum[favoriteTypei−1]+1, sum[favoriteTypei]]。最后判断 query 区间和累积吃糖果数的区间是否有重叠即可。如果重叠即输出 true。
- 判断两个区间是否重合，情况有好几种：内包含，全包含，半包含等等。没有交集的情况比较少，所以可以用排除法。对于区间 [x1, y1] 以及 [x2, y2]，它们没有交集当且仅当 x1 > y2 或者 y1 < x2。

## 代码

```go
package leetcode

func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
	prefixSum := make([]int, n)
	prefixSum[0] = candiesCount[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + candiesCount[i]
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		favoriteType, favoriteDay, dailyCap := q[0], q[1], q[2]
		x1 := favoriteDay + 1
		y1 := (favoriteDay + 1) * dailyCap
		x2 := 1
		if favoriteType > 0 {
			x2 = prefixSum[favoriteType-1] + 1
		}
		y2 := prefixSum[favoriteType]
		res[i] = !(x1 > y2 || y1 < x2)
	}
	return res
}
```