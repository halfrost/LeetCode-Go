# [862. Shortest Subarray with Sum at Least K](https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/)



## 题目

Return the **length** of the shortest, non-empty, contiguous subarray of `A` with sum at least `K`.

If there is no non-empty subarray with sum at least `K`, return `-1`.

**Example 1**:

```
Input: A = [1], K = 1
Output: 1
```

**Example 2**:

```
Input: A = [1,2], K = 4
Output: -1
```

**Example 3**:

```
Input: A = [2,-1,2], K = 3
Output: 3
```

**Note**:

1. `1 <= A.length <= 50000`
2. `-10 ^ 5 <= A[i] <= 10 ^ 5`
3. `1 <= K <= 10 ^ 9`

## 题目大意

返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。如果没有和至少为 K 的非空子数组，返回 -1 。

提示：

- 1 <= A.length <= 50000
- -10 ^ 5 <= A[i] <= 10 ^ 5
- 1 <= K <= 10 ^ 9


## 解题思路

- 给出一个数组，要求找出一个最短的，非空的，连续的子序列且累加和至少为 k 。
- 由于给的数组里面可能存在负数，所以子数组的累加和不会随着数组的长度增加而增加。题目中要求区间和，理所应当需要利用 `prefixSum` 前缀和，先计算出 `prefixSum` 前缀和。
- 简化一下题目的要求，即能否找到 `prefixSum[y] - prefixSum[x] ≥ K` ，且 `y - x` 的差值最小。如果固定的 `y`，那么对于 `x`，`x` 越大，`y - x` 的差值就越小(因为 `x` 越逼近 `y`)。所以想求区间 `[x, y]` 的最短距离，需要保证 `y` 尽量小，`x` 尽量大，这样 `[x, y]` 区间距离最小。那么有以下 2 点“常识”一定成立：
    1. 如果 `x1 < x2` ，并且 `prefixSum[x2] ≤ prefixSum[x1]`，说明结果一定不会取 `x1`。因为如果 `prefixSum[x1] ≤ prefixSum[y] - k`，那么  `prefixSum[x2] ≤ prefixSum[x1] ≤ prefixSum[y] - k`，`x2` 也能满足题意，并且 `x2` 比 `x1` 更加接近 `y`，最优解一定优先考虑 `x2`。
    2. 在确定了 `x` 以后，以后就不用再考虑 `x` 了，因为如果 `y2 > y1`，且 `y2` 的时候取 `x` 还是一样的，那么算距离的话，`y2 - x` 显然大于 `y1 - x`，这样的话肯定不会是最短的距离。
- 从上面这两个常识来看，可以用双端队列 `deque` 来处理 `prefixSum`。`deque` 中存储的是递增的 `x` 下标，为了满足常识一。从双端队列的开头开始遍历，假如区间和之差大于等于 `K`，就移除队首元素并更新结果 `res`。队首移除元素，直到不满足 `prefixSum[i]-prefixSum[deque[0]] >= K` 这一不等式，是为了满足常识二。之后的循环是此题的精髓，从双端队列的末尾开始往前遍历，假如当前区间和 `prefixSum[i]` 小于等于队列末尾的区间和，则移除队列末尾元素。为什么这样处理呢？因为若数组都是正数，那么长度越长，区间和一定越大，则 `prefixSum[i]` 一定大于所有双端队列中的区间和，但由于可能存在负数，从而使得长度变长，区间总和反而减少了，之前的区间和之差值都没有大于等于 `K`(< K)，现在的更不可能大于等于 `K`，这个结束位置可以直接淘汰，不用进行计算。循环结束后将当前位置加入双端队列即可。遇到新下标在队尾移除若干元素，这一行为，也是为了满足常识一。
- 由于所有下标都只进队列一次，也最多 pop 出去一次，所以时间复杂度 O(n)，空间复杂度 O(n)。

## 代码

```go
func shortestSubarray(A []int, K int) int {
	res, prefixSum := len(A)+1, make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		prefixSum[i+1] = prefixSum[i] + A[i]
	}
	// deque 中保存递增的 prefixSum 下标
	deque := []int{}
	for i := range prefixSum {
		// 下面这个循环希望能找到 [deque[0], i] 区间内累加和 >= K，如果找到了就更新答案
		for len(deque) > 0 && prefixSum[i]-prefixSum[deque[0]] >= K {
			length := i - deque[0]
			if res > length {
				res = length
			}
			// 找到第一个 deque[0] 能满足条件以后，就移除它，因为它是最短长度的子序列了
			deque = deque[1:]
		}
		// 下面这个循环希望能保证 prefixSum[deque[i]] 递增
		for len(deque) > 0 && prefixSum[i] <= prefixSum[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	if res <= len(A) {
		return res
	}
	return -1
}
```