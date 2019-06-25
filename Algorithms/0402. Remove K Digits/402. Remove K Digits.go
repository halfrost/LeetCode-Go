package leetcode

func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	res := []byte{}
	for i := 0; i < len(num); i++ {
		c := num[i]
		for k > 0 && len(res) > 0 && c < res[len(res)-1] {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, c)
	}
	res = res[:len(res)-k]

	// trim leading zeros
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return string(res)
}
