package leetcode

func equalSubstring(s string, t string, maxCost int) int {
	left, right, res := 0, -1, 0
	for left < len(s) {
		if right+1 < len(s) && maxCost-abs(int(s[right+1]-'a')-int(t[right+1]-'a')) >= 0 {
			right++
			maxCost -= abs(int(s[right]-'a') - int(t[right]-'a'))
		} else {
			res = max(res, right-left+1)
			maxCost += abs(int(s[left]-'a') - int(t[left]-'a'))
			left++
		}
	}
	return res
}
