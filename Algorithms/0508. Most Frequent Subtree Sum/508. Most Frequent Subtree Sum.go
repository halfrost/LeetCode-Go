package leetcode

import (
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法一 维护最大频次，不用排序
func findFrequentTreeSum(root *TreeNode) []int {
	memo := make(map[int]int)
	collectSum(root, memo)
	res := []int{}
	most := 0
	for key, val := range memo {
		if most == val {
			res = append(res, key)
		} else if most < val {
			most = val
			res = []int{key}
		}
	}
	return res
}

func collectSum(root *TreeNode, memo map[int]int) int {
	if root == nil {
		return 0
	}
	sum := root.Val + collectSum(root.Left, memo) + collectSum(root.Right, memo)
	if v, ok := memo[sum]; ok {
		memo[sum] = v + 1
	} else {
		memo[sum] = 1
	}
	return sum
}

// 解法二 求出所有和再排序
func findFrequentTreeSum1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	freMap, freList, reFreMap := map[int]int{}, []int{}, map[int][]int{}
	findTreeSum(root, freMap)
	for k, v := range freMap {
		tmp := reFreMap[v]
		tmp = append(tmp, k)
		reFreMap[v] = tmp
	}
	for k := range reFreMap {
		freList = append(freList, k)
	}
	sort.Ints(freList)
	return reFreMap[freList[len(freList)-1]]
}

func findTreeSum(root *TreeNode, fre map[int]int) int {
	if root == nil {
		return 0
	}
	if root != nil && root.Left == nil && root.Right == nil {
		fre[root.Val]++
		return root.Val
	}
	val := findTreeSum(root.Left, fre) + findTreeSum(root.Right, fre) + root.Val
	fre[val]++
	return val
}
