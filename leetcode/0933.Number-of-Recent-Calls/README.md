# [933. Number of Recent Calls](https://leetcode.com/problems/number-of-recent-calls/)



## 题目

Write a class `RecentCounter` to count recent requests.

It has only one method: `ping(int t)`, where t represents some time in milliseconds.

Return the number of `ping`s that have been made from 3000 milliseconds ago until now.

Any ping with time in `[t - 3000, t]` will count, including the current ping.

It is guaranteed that every call to `ping` uses a strictly larger value of `t` than before.

**Example 1**:

```
Input: inputs = ["RecentCounter","ping","ping","ping","ping"], inputs = [[],[1],[100],[3001],[3002]]
Output: [null,1,2,3,3]
```

**Note**:

1. Each test case will have at most `10000` calls to `ping`.
2. Each test case will call `ping` with strictly increasing values of `t`.
3. Each call to ping will have `1 <= t <= 10^9`.


## 题目大意

写一个 RecentCounter 类来计算最近的请求。它只有一个方法：ping(int t)，其中 t 代表以毫秒为单位的某个时间。返回从 3000 毫秒前到现在的 ping 数。任何处于 [t - 3000, t] 时间范围之内的 ping 都将会被计算在内，包括当前（指 t 时刻）的 ping。保证每次对 ping 的调用都使用比之前更大的 t 值。
 
提示：

- 每个测试用例最多调用 10000 次 ping。
- 每个测试用例会使用严格递增的 t 值来调用 ping。
- 每次调用 ping 都有 1 <= t <= 10^9。


## 解题思路

- 要求设计一个类，可以用 `ping(t)` 的方法，计算 [t-3000, t] 区间内的 ping 数。t 是毫秒。
- 这一题比较简单，`ping()` 方法用二分搜索即可。

## 代码

```go
type RecentCounter struct {
	list []int
}

func Constructor933() RecentCounter {
	return RecentCounter{
		list: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.list = append(this.list, t)
	index := sort.Search(len(this.list), func(i int) bool { return this.list[i] >= t-3000 })
	if index < 0 {
		index = -index - 1
	}
	return len(this.list) - index
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
```