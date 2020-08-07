package leetcode

import "strings"

func isValidSerialization(preorder string) bool {
	nodes, diff := strings.Split(preorder, ","), 1
	for _, node := range nodes {
		diff--
		if diff < 0 {
			return false
		}
		if node != "#" {
			diff += 2
		}
	}
	return diff == 0
}
