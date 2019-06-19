package leetcode

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int, len(A)*len(B))
	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}
	ret := 0
	for _, c := range C {
		for _, d := range D {
			ret += m[0-c-d]
		}
	}

	return ret
}
