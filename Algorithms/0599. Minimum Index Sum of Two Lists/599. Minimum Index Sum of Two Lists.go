package leetcode

func findRestaurant(list1 []string, list2 []string) []string {
	m, ans := make(map[string]int, len(list1)), []string{}
	for i, r := range list1 {
		m[r] = i
	}
	for j, r := range list2 {
		if _, ok := m[r]; ok {
			m[r] += j
			if len(ans) == 0 || m[r] == m[ans[0]] {
				ans = append(ans, r)
			} else if m[r] < m[ans[0]] {
				ans = []string{r}
			}
		}
	}
	return ans
}
