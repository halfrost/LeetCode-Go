# [352. Data Stream as Disjoint Intervals](https://leetcode.com/problems/data-stream-as-disjoint-intervals/)


## 题目

Given a data stream input of non-negative integers a1, a2, ..., an, summarize the numbers seen so far as a list of disjoint intervals.

Implement the SummaryRanges class:

   - SummaryRanges() Initializes the object with an empty stream.
   - void addNum(int val) Adds the integer val to the stream. 
   - int[][] getIntervals() Returns a summary of the integers in the stream currently as a list of disjoint intervals [starti, endi].

**Example 1:**

    Input
    ["SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"]
    [[], [1], [], [3], [], [7], [], [2], [], [6], []]
    Output
    [null, null, [[1, 1]], null, [[1, 1], [3, 3]], null, [[1, 1], [3, 3], [7, 7]], null, [[1, 3], [7, 7]], null, [[1, 3], [6, 7]]]

    Explanation
    SummaryRanges summaryRanges = new SummaryRanges();
    summaryRanges.addNum(1);      // arr = [1]
    summaryRanges.getIntervals(); // return [[1, 1]]
    summaryRanges.addNum(3);      // arr = [1, 3]
    summaryRanges.getIntervals(); // return [[1, 1], [3, 3]]
    summaryRanges.addNum(7);      // arr = [1, 3, 7]
    summaryRanges.getIntervals(); // return [[1, 1], [3, 3], [7, 7]]
    summaryRanges.addNum(2);      // arr = [1, 2, 3, 7]
    summaryRanges.getIntervals(); // return [[1, 3], [7, 7]]
    summaryRanges.addNum(6);      // arr = [1, 2, 3, 6, 7]
    summaryRanges.getIntervals(); // return [[1, 3], [6, 7]]

**Constraints**

   - 0 <= val <= 10000
   - At most 3 * 10000 calls will be made to addNum and getIntervals.

## 题目大意

给你一个由非负整数a1, a2, ..., an 组成的数据流输入，请你将到目前为止看到的数字总结为不相交的区间列表。

实现 SummaryRanges 类：

   - SummaryRanges() 使用一个空数据流初始化对象。
   - void addNum(int val) 向数据流中加入整数 val 。
   - int[][] getIntervals() 以不相交区间[starti, endi] 的列表形式返回对数据流中整数的总结

## 解题思路

- 使用字典过滤掉重复的数字
- 把过滤后的数字放到nums中,并进行排序 
- 使用nums构建不重复的区间

## 代码

```go
package leetcode

import "sort"

type SummaryRanges struct {
	nums []int
	mp map[int]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		nums: []int{},
		mp : map[int]int{},
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
		if this.nums[i] == end + 1 {
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
```