package leetcode

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	for _, v := range words {
		if exist(board, v) {
			res = append(res, v)
		}
	}
	return res
}
