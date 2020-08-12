package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for a, b := 0, len(A[i])-1; a < b; a, b = a+1, b-1 {
			A[i][a], A[i][b] = A[i][b], A[i][a]
		}
		for a := 0; a < len(A[i]); a++ {
			A[i][a] = (A[i][a] + 1) % 2
		}
	}
	return A
}
