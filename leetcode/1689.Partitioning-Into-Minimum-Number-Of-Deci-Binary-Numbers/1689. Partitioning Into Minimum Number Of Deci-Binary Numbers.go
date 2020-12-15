package leetcode

func minPartitions(n string) int {
	res := 0
	for i := 0; i < len(n); i++ {
		if int(n[i]-'0') > res {
			res = int(n[i] - '0')
		}
	}
	return res
}
