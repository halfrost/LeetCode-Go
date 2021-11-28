package leetcode

import "sort"

type SummaryRanges struct {
	nums []int
	mp   map[int]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		nums: []int{},
		mp:   map[int]int{},
	}
}

func (this *SummaryRanges) AddNum(val int) {
	if _, ok := this.mp[val]; !ok {
		this.mp[val] = 1
		this.nums = append(this.nums, val)
	}
	sort.Ints(this.nums)
}

func (this *SummaryRanges) GetIntervals() [][]int {
	n := len(this.nums)
	var ans [][]int
	if n == 0 {
		return ans
	}
	if n == 1 {
		ans = append(ans, []int{this.nums[0], this.nums[0]})
		return ans
	}
	start, end := this.nums[0], this.nums[0]
	ans = append(ans, []int{start, end})
	index := 0
	for i := 1; i < n; i++ {
		if this.nums[i] == end+1 {
			end = this.nums[i]
			ans[index][1] = end
		} else {
			start, end = this.nums[i], this.nums[i]
			ans = append(ans, []int{start, end})
			index++
		}
	}
	return ans
}
