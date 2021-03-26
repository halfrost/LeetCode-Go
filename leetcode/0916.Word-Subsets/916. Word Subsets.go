package leetcode

func wordSubsets(A []string, B []string) []string {
	var counter [26]int
	for _, b := range B {
		var m [26]int
		for _, c := range b {
			j := c - 'a'
			m[j]++
		}
		for i := 0; i < 26; i++ {
			if m[i] > counter[i] {
				counter[i] = m[i]
			}
		}
	}
	var res []string
	for _, a := range A {
		var m [26]int
		for _, c := range a {
			j := c - 'a'
			m[j]++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if m[i] < counter[i] {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, a)
		}
	}
	return res
}
