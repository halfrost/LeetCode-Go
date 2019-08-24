package leetcode

// 解法一 滑动窗口优化版
func maxSatisfied(customers []int, grumpy []int, X int) int {
	customer0, customer1, maxCustomer1, left, right := 0, 0, 0, 0, 0
	for ; right < len(customers); right++ {
		if grumpy[right] == 0 {
			customer0 += customers[right]
		} else {
			customer1 += customers[right]
			for right-left+1 > X {
				if grumpy[left] == 1 {
					customer1 -= customers[left]
				}
				left++
			}
			if customer1 > maxCustomer1 {
				maxCustomer1 = customer1
			}
		}
	}
	return maxCustomer1 + customer0
}

// 解法二 滑动窗口暴力版
func maxSatisfied1(customers []int, grumpy []int, X int) int {
	left, right, res := 0, -1, 0
	for left < len(customers) {
		if right+1 < len(customers) && right-left < X-1 {
			right++
		} else {
			if right-left+1 == X {
				res = max(res, sumSatisfied(customers, grumpy, left, right))
			}
			left++
		}
	}
	return res
}

func sumSatisfied(customers []int, grumpy []int, start, end int) int {
	sum := 0
	for i := 0; i < len(customers); i++ {
		if i < start || i > end {
			if grumpy[i] == 0 {
				sum += customers[i]
			}
		} else {
			sum += customers[i]
		}
	}
	return sum
}
