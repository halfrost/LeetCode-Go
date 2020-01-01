package leetcode

import (
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
	if len(S) == 0 {
		return &TreeNode{}
	}
	root, index, level := &TreeNode{}, 0, 0
	cur := root
	dfsBuildPreorderTree(S, &index, &level, cur)
	return root.Right
}

func dfsBuildPreorderTree(S string, index, level *int, cur *TreeNode) (newIndex *int) {
	if *index == len(S) {
		return index
	}
	if *index == 0 && *level == 0 {
		i := 0
		for i = *index; i < len(S); i++ {
			if !isDigital(S[i]) {
				break
			}
		}
		num, _ := strconv.Atoi(S[*index:i])
		tmp := &TreeNode{Val: num, Left: nil, Right: nil}
		cur.Right = tmp
		nLevel := *level + 1
		index = dfsBuildPreorderTree(S, &i, &nLevel, tmp)
		index = dfsBuildPreorderTree(S, index, &nLevel, tmp)
	}
	i := 0
	for i = *index; i < len(S); i++ {
		if isDigital(S[i]) {
			break
		}
	}
	if *level == i-*index {
		j := 0
		for j = i; j < len(S); j++ {
			if !isDigital(S[j]) {
				break
			}
		}
		num, _ := strconv.Atoi(S[i:j])
		tmp := &TreeNode{Val: num, Left: nil, Right: nil}
		if cur.Left == nil {
			cur.Left = tmp
			nLevel := *level + 1
			index = dfsBuildPreorderTree(S, &j, &nLevel, tmp)
			index = dfsBuildPreorderTree(S, index, level, cur)
		} else if cur.Right == nil {
			cur.Right = tmp
			nLevel := *level + 1
			index = dfsBuildPreorderTree(S, &j, &nLevel, tmp)
			index = dfsBuildPreorderTree(S, index, level, cur)
		}
	}
	return index
}
