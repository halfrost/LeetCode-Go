package leetcode

import "sort"

// 解法一 利用系统函数 sort + 二分搜索
func findRightInterval(intervals [][]int) []int {
	intervalList := make(intervalList, len(intervals))
	// 转换成 interval 类型
	for i, v := range intervals {
		intervalList[i] = interval{interval: v, index: i}
	}
	sort.Sort(intervalList)
	out := make([]int, len(intervalList))
	for i := 0; i < len(intervalList); i++ {
		index := sort.Search(len(intervalList), func(p int) bool { return intervalList[p].interval[0] >= intervalList[i].interval[1] })
		if index == len(intervalList) {
			out[intervalList[i].index] = -1
		} else {
			out[intervalList[i].index] = intervalList[index].index
		}
	}
	return out
}

type interval struct {
	interval []int
	index    int
}

type intervalList []interval

func (in intervalList) Len() int { return len(in) }
func (in intervalList) Less(i, j int) bool {
	return in[i].interval[0] <= in[j].interval[0]
}
func (in intervalList) Swap(i, j int) { in[i], in[j] = in[j], in[i] }

// 解法二 手撸 sort + 二分搜索
func findRightInterval1(intervals [][]int) []int {
	if len(intervals) == 0 {
		return []int{}
	}
	intervalsList, res, intervalMap := []Interval{}, []int{}, map[Interval]int{}
	for k, v := range intervals {
		intervalsList = append(intervalsList, Interval{Start: v[0], End: v[1]})
		intervalMap[Interval{Start: v[0], End: v[1]}] = k
	}
	quickSort(intervalsList, 0, len(intervalsList)-1)
	for _, v := range intervals {
		tmp := searchFirstGreaterInterval(intervalsList, v[1])
		if tmp > 0 {
			tmp = intervalMap[intervalsList[tmp]]
		}
		res = append(res, tmp)
	}
	return res
}

func searchFirstGreaterInterval(nums []Interval, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid].Start >= target {
			if (mid == 0) || (nums[mid-1].Start < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
