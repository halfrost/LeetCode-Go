# [638. Shopping Offers](https://leetcode.com/problems/shopping-offers/)



## 题目

In LeetCode Store, there are some kinds of items to sell. Each item has a price.

However, there are some special offers, and a special offer consists of one or more different kinds of items with a sale price.

You are given the each item's price, a set of special offers, and the number we need to buy for each item. The job is to output the lowest price you have to pay for **exactly** certain items as given, where you could make optimal use of the special offers.

Each special offer is represented in the form of an array, the last number represents the price you need to pay for this special offer, other numbers represents how many specific items you could get if you buy this offer.

You could use any of special offers as many times as you want.

**Example 1**:

```
Input: [2,5], [[3,0,5],[1,2,10]], [3,2]
Output: 14
Explanation: 
There are two kinds of items, A and B. Their prices are $2 and $5 respectively. 
In special offer 1, you can pay $5 for 3A and 0B
In special offer 2, you can pay $10 for 1A and 2B. 
You need to buy 3A and 2B, so you may pay $10 for 1A and 2B (special offer #2), and $4 for 2A.
```

**Example 2**:

```
Input: [2,3,4], [[1,1,0,4],[2,2,1,9]], [1,2,1]
Output: 11
Explanation: 
The price of A is $2, and $3 for B, $4 for C. 
You may pay $4 for 1A and 1B, and $9 for 2A ,2B and 1C. 
You need to buy 1A ,2B and 1C, so you may pay $4 for 1A and 1B (special offer #1), and $3 for 1B, $4 for 1C. 
You cannot add more items, though only $9 for 2A ,2B and 1C.
```

**Note**:

1. There are at most 6 kinds of items, 100 special offers.
2. For each item, you need to buy at most 6 of them.
3. You are **not** allowed to buy more items than you want, even if that would lower the overall price.


## 题目大意

在 LeetCode 商店中， 有许多在售的物品。然而，也有一些大礼包，每个大礼包以优惠的价格捆绑销售一组物品。

现给定每个物品的价格，每个大礼包包含物品的清单，以及待购物品清单。请输出确切完成待购清单的最低花费。每个大礼包的由一个数组中的一组数据描述，最后一个数字代表大礼包的价格，其他数字分别表示内含的其他种类物品的数量。任意大礼包可无限次购买。

例子 1:

```
输入: [2,5], [[3,0,5],[1,2,10]], [3,2]
输出: 14
解释: 
有A和B两种物品，价格分别为¥2和¥5。
大礼包1，你可以以¥5的价格购买3A和0B。
大礼包2， 你可以以¥10的价格购买1A和2B。
你需要购买3个A和2个B， 所以你付了¥10购买了1A和2B（大礼包2），以及¥4购买2A。

```
例子 2:

```
输入: [2,3,4], [[1,1,0,4],[2,2,1,9]], [1,2,1]
输出: 11
解释: 
A，B，C的价格分别为¥2，¥3，¥4.
你可以用¥4购买1A和1B，也可以用¥9购买2A，2B和1C。
你需要买1A，2B和1C，所以你付了¥4买了1A和1B（大礼包1），以及¥3购买1B， ¥4购买1C。
你不可以购买超出待购清单的物品，尽管购买大礼包2更加便宜。
```

说明:

- 最多6种物品， 100种大礼包。
- 每种物品，你最多只需要购买6个。
- 你不可以购买超出待购清单的物品，即使更便宜。


## 解题思路

- 给出 3 个数组，3 个数组分别代表的意义是在售的商品价格，多个礼包以及礼包内每个商品的数量和总价，购物清单上需要购买每个商品的数量。问购买清单上的所有商品所需的最低花费。
- 这一题可以用 DFS 暴力解题，也可以用 DP。笔者这题先用 DFS 来解答。设当前搜索到的状态为 `shopping(price, special, needs)`，其中 `price` 和 `special` 为题目中所述的物品的单价和捆绑销售的大礼包，而 `needs` 为当前需要的每种物品的数量。针对于每个商品，可以有 3 种购买规则，第一种，选礼包里面的第一个优惠购买，第二种，不选当前礼包优惠，选下一个优惠进行购买，第三种，不使用优惠，直接购买。这样就对应了 3 种 DFS 的方向。具体见代码。如果选择了礼包优惠，那么递归到下一层，`need` 需要对应减少礼包里面的数量，最终金额累加。当所有情况遍历完以后，可以返回出最小花费。
- 这一题需要注意的剪枝情况：是否需要购买礼包。题目中要求了，不能购买超过清单上数量的商品，即使价格便宜，也不行。例如可以买 n 个礼包 A，但是最终商品数量超过了清单上的商品，这种购买方式是不行的。所以需要先判断当前递归中，满足 `need` 和 `price` 条件的，能否使用礼包。这里包含 2 种情况，一种是当前商品已经满足清单个数了，不需要再买了；还有一种情况是已经超过清单数量了，那这种情况需要立即返回，当前这种购买方式不合题意。

## 代码

```go
func shoppingOffers(price []int, special [][]int, needs []int) int {
	res := -1
	dfsShoppingOffers(price, special, needs, 0, &res)
	return res
}

func dfsShoppingOffers(price []int, special [][]int, needs []int, pay int, res *int) {
	noNeeds := true
	// 剪枝
	for _, need := range needs {
		if need < 0 {
			return
		}
		if need != 0 {
			noNeeds = false
		}
	}
	if len(special) == 0 || noNeeds {
		for i, p := range price {
			pay += (p * needs[i])
		}
		if pay < *res || *res == -1 {
			*res = pay
		}
		return
	}
	newNeeds := make([]int, len(needs))
	copy(newNeeds, needs)
	for i, n := range newNeeds {
		newNeeds[i] = n - special[0][i]
	}
	dfsShoppingOffers(price, special, newNeeds, pay+special[0][len(price)], res)
	dfsShoppingOffers(price, special[1:], newNeeds, pay+special[0][len(price)], res)
	dfsShoppingOffers(price, special[1:], needs, pay, res)
}
```