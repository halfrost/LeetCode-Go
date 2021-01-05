# [830. Positions of Large Groups](https://leetcode.com/problems/positions-of-large-groups/)


## 题目

In a string `s` of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like `s = "abbxxxxzyy"` has the groups `"a"`, `"bb"`, `"xxxx"`, `"z"`, and `"yy"`.

A group is identified by an interval `[start, end]`, where `start` and `end` denote the start and end indices (inclusive) of the group. In the above example, `"xxxx"` has the interval `[3,6]`.

A group is considered **large** if it has 3 or more characters.

Return *the intervals of every **large** group sorted in **increasing order by start index***.

**Example 1:**

```
Input: s = "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the only large group with start index 3 and end index 6.
```

**Example 2:**

```
Input: s = "abc"
Output: []
Explanation: We have groups "a", "b", and "c", none of which are large groups.
```

**Example 3:**

```
Input: s = "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]
Explanation: The large groups are "ddd", "eeee", and "bbb".
```

**Example 4:**

```
Input: s = "aba"
Output: []
```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` contains lower-case English letters only.

## 题目大意

在一个由小写字母构成的字符串 s 中，包含由一些连续的相同字符所构成的分组。例如，在字符串 s = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。分组可以用区间 [start, end] 表示，其中 start 和 end 分别表示该分组的起始和终止位置的下标。上例中的 "xxxx" 分组用区间表示为 [3,6] 。我们称所有包含大于或等于三个连续字符的分组为 较大分组 。

找到每一个 较大分组 的区间，按起始位置下标递增顺序排序后，返回结果。

## 解题思路

- 简单题。利用滑动窗口的思想，先扩大窗口的右边界，找到能相同字母且能到达的最右边。记录左右边界。再将窗口的左边界移动到上一次的右边界处。以此类推，重复扩大窗口的右边界，直至扫完整个字符串。最终所有满足题意的较大分组区间都在数组中了。

## 代码

```go
package leetcode

func largeGroupPositions(S string) [][]int {
	res, end := [][]int{}, 0
	for end < len(S) {
		start, str := end, S[end]
		for end < len(S) && S[end] == str {
			end++
		}
		if end-start >= 3 {
			res = append(res, []int{start, end - 1})
		}
	}
	return res
}
```