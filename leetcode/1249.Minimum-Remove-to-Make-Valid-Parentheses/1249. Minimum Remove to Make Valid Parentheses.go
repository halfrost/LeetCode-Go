package leetcode

func minRemoveToMakeValid(s string) string {
	res, opens := []byte{}, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			opens++
		} else if s[i] == ')' {
			if opens == 0 {
				continue
			}
			opens--
		}
		res = append(res, s[i])
	}
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == '(' && opens > 0 {
			opens--
			res = append(res[:i], res[i+1:]...)
		}
	}
	return string(res)
}
