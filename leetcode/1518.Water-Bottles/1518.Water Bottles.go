package leetcode

func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles < numExchange {
		return numBottles
	}
	quotient := numBottles / numExchange
	reminder := numBottles % numExchange
	ans := numBottles + quotient
	for quotient+reminder >= numExchange {
		quotient, reminder = (quotient+reminder)/numExchange, (quotient+reminder)%numExchange
		ans += quotient
	}
	return ans
}
