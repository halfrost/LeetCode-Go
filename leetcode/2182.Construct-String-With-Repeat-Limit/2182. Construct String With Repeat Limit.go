package leetcode

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[int(c-'a')]++
	}
	var ns []byte
	for i := 25; i >= 0; {
		k := i - 1
		for cnt[i] > 0 {
			for j := 0; j < min(cnt[i], repeatLimit); j++ {
				ns = append(ns, byte(i)+'a')
			}
			cnt[i] -= repeatLimit
			if cnt[i] > 0 {
				for ; k >= 0 && cnt[k] == 0; k-- {
				}
				if k < 0 {
					break
				} else {
					ns = append(ns, byte(k)+'a')
					cnt[k]--
				}
			}
		}
		i = k
	}
	return string(ns)
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
