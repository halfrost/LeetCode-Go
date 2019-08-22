package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 map，时间复杂度 O(n)
func longestConsecutive(nums []int) int {
	res, numMap := 0, map[int]int{}
	for _, num := range nums {
		if numMap[num] == 0 {
			left, right, sum := 0, 0, 0
			if numMap[num-1] > 0 {
				left = numMap[num-1]
			} else {
				left = 0
			}
			if numMap[num+1] > 0 {
				right = numMap[num+1]
			} else {
				right = 0
			}
			// sum: length of the sequence n is in
			sum = left + right + 1
			numMap[num] = sum
			// keep track of the max length
			res = max(res, sum)
			// extend the length to the boundary(s) of the sequence
			// will do nothing if n has no neighbors
			numMap[num-left] = sum
			numMap[num+right] = sum
		} else {
			continue
		}
	}
	return res
}

// 解法二 并查集
func longestConsecutive1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap, countMap, lcs, uf := map[int]int{}, map[int]int{}, 0, template.UnionFind{}
	uf.Init(len(nums))
	for i := 0; i < len(nums); i++ {
		countMap[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			continue
		}
		numMap[nums[i]] = i
		if _, ok := numMap[nums[i]+1]; ok {
			uf.Union(i, numMap[nums[i]+1])
		}
		if _, ok := numMap[nums[i]-1]; ok {
			uf.Union(i, numMap[nums[i]-1])
		}
	}
	for key := range countMap {
		parent := uf.Find(key)
		if parent != key {
			countMap[parent]++
		}
		if countMap[parent] > lcs {
			lcs = countMap[parent]
		}
	}
	return lcs
}

// 解法三 暴力解法，时间复杂度 O(n^2)
func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap, length, tmp, lcs := map[int]bool{}, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = true
	}
	for key := range numMap {
		if !numMap[key-1] && !numMap[key+1] {
			delete(numMap, key)
		}
	}
	if len(numMap) == 0 {
		return 1
	}
	for key := range numMap {
		if !numMap[key-1] && numMap[key+1] {
			length, tmp = 1, key+1
			for numMap[tmp] {
				length++
				tmp++
			}
			lcs = max(lcs, length)
		}
	}
	return max(lcs, length)
}
