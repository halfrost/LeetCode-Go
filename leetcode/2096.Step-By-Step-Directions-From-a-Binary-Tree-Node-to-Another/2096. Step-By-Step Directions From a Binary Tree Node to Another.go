package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getDirections(root *TreeNode, startValue int, destValue int) string {
	sPath, dPath := make([]byte, 0), make([]byte, 0)
	findPath(root, startValue, &sPath)
	findPath(root, destValue, &dPath)
	size, i := min(len(sPath), len(dPath)), 0
	for i < size {
		if sPath[len(sPath)-1-i] == dPath[len(dPath)-1-i] {
			i++
		} else {
			break
		}
	}
	sPath = sPath[:len(sPath)-i]
	replace(sPath)
	dPath = dPath[:len(dPath)-i]
	reverse(dPath)
	sPath = append(sPath, dPath...)
	return string(sPath)
}

func findPath(root *TreeNode, value int, path *[]byte) bool {
	if root.Val == value {
		return true
	}

	if root.Left != nil && findPath(root.Left, value, path) {
		*path = append(*path, 'L')
		return true
	}

	if root.Right != nil && findPath(root.Right, value, path) {
		*path = append(*path, 'R')
		return true
	}

	return false
}

func reverse(path []byte) {
	left, right := 0, len(path)-1
	for left < right {
		path[left], path[right] = path[right], path[left]
		left++
		right--
	}
}

func replace(path []byte) {
	for i := 0; i < len(path); i++ {
		path[i] = 'U'
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
