package leetcode

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	minRad := 0
	sort.Ints(heaters)
	for _, house := range houses {
		// 遍历房子的坐标，维护 heaters 的最小半径
		heater := findClosestHeater(house, heaters)
		rad := heater - house
		if rad < 0 {
			rad = -rad
		}
		if rad > minRad {
			minRad = rad
		}
	}
	return minRad

}

// 二分搜索
func findClosestHeater(pos int, heaters []int) int {
	low, high := 0, len(heaters)-1
	if pos < heaters[low] {
		return heaters[low]
	}
	if pos > heaters[high] {
		return heaters[high]
	}
	for low <= high {
		mid := low + (high-low)>>1
		if pos == heaters[mid] {
			return heaters[mid]
		} else if pos < heaters[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	// 判断距离两边的 heaters 哪个更近
	if pos-heaters[high] < heaters[low]-pos {
		return heaters[high]
	}
	return heaters[low]
}

// 解法二 暴力搜索
func findRadius1(houses []int, heaters []int) int {
	res := 0
	for i := 0; i < len(houses); i++ {
		dis := math.MaxInt64
		for j := 0; j < len(heaters); j++ {
			dis = min(dis, abs(houses[i]-heaters[j]))
		}
		res = max(res, dis)
	}
	return res
}
