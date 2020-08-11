package leetcode

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	cur, res := 0, []int{}
	for _, v := range A {
		if v%2 == 0 {
			cur += v
		}
	}
	for _, q := range queries {
		if A[q[1]]%2 == 0 {
			cur -= A[q[1]]
		}
		A[q[1]] += q[0]
		if A[q[1]]%2 == 0 {
			cur += A[q[1]]
		}
		res = append(res, cur)
	}
	return res
}
