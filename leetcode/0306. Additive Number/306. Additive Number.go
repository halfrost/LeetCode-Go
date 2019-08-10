package leetcode

import (
	"strconv"
	"strings"
)

// This function controls various combinations as starting points
func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	for firstEnd := 0; firstEnd < len(num)/2; firstEnd++ {
		if num[0] == '0' && firstEnd > 0 {
			break
		}
		first, _ := strconv.Atoi(num[:firstEnd+1])
		for secondEnd := firstEnd + 1; max(firstEnd, secondEnd-firstEnd) <= len(num)-secondEnd; secondEnd++ {
			if num[firstEnd+1] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			second, _ := strconv.Atoi(num[firstEnd+1 : secondEnd+1])
			if recursiveCheck(num, first, second, secondEnd+1) {
				return true
			}
		}
	}
	return false
}

//Propagate for rest of the string
func recursiveCheck(num string, x1 int, x2 int, left int) bool {
	if left == len(num) {
		return true
	}
	if strings.HasPrefix(num[left:], strconv.Itoa(x1+x2)) {
		return recursiveCheck(num, x2, x1+x2, left+len(strconv.Itoa(x1+x2)))
	}
	return false
}
