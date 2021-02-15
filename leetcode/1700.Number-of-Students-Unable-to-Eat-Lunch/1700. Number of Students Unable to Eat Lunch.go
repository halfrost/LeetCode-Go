package leetcode

func countStudents(students []int, sandwiches []int) int {
	tmp, n, i := [2]int{}, len(students), 0
	for _, v := range students {
		tmp[v]++
	}
	for i < n && tmp[sandwiches[i]] > 0 {
		tmp[sandwiches[i]]--
		i++
	}
	return n - i
}
