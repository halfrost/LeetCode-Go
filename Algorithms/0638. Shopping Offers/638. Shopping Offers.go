package leetcode

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
