package leetcode

func fairCandySwap(A []int, B []int) []int {
	hDiff, aMap := diff(A, B)/2, make(map[int]int, len(A))
	for _, a := range A {
		aMap[a] = a
	}
	for _, b := range B {
		if a, ok := aMap[hDiff+b]; ok {
			return []int{a, b}
		}
	}
	return nil
}

func diff(A []int, B []int) int {
	diff, maxLen := 0, max(len(A), len(B))
	for i := 0; i < maxLen; i++ {
		if i < len(A) {
			diff += A[i]
		}
		if i < len(B) {
			diff -= B[i]
		}
	}
	return diff
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
