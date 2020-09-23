# [1512. Number of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/)

## 题目

Given an array of integers `nums`.

A pair `(i,j)` is called good if `nums[i] == nums[j]` and `i < j`.

Return the number of good pairs.

**Example 1**:

```
Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

```

**Example 2**:

```
Input: nums = [1,1,1,1]
Output: 6
Explanation: Each pair in the array are good.

```

**Example 3**:

```
Input: nums = [1,2,3]
Output: 0

```

**Constraints**:

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 100`

```go
func numIdenticalPairs(nums []int) int {
	total := 0
	for x := 0; x < len(nums); x++ {
		for y := x + 1; y < len(nums); y++ {
			if nums[x] == nums[y] {
				total++
			}
		}
	}

	return total
}


```