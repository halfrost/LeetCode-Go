package leetcode

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	b1, b2, tmp := []byte(num1), []byte(num2), make([]int, len(num1)+len(num2))
	for i := 0; i < len(b1); i++ {
		for j := 0; j < len(b2); j++ {
			tmp[i+j+1] += int(b1[i]-'0') * int(b2[j]-'0')
		}
	}
	for i := len(tmp) - 1; i > 0; i-- {
		tmp[i-1] += tmp[i] / 10
		tmp[i] = tmp[i] % 10
	}
	if tmp[0] == 0 {
		tmp = tmp[1:]
	}
	res := make([]byte, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[i] = '0' + byte(tmp[i])
	}
	return string(res)
}
