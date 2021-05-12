package leetcode

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr))
	xors[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		xors[i] = arr[i] ^ xors[i-1]
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = xors[q[1]]
		if q[0] > 0 {
			res[i] ^= xors[q[0]-1]
		}
	}
	return res
}
