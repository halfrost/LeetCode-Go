package leetcode

func numMatchingSubseq(s string, words []string) int {
	hash, res := make([][]string, 26), 0
	for _, w := range words {
		hash[int(w[0]-'a')] = append(hash[int(w[0]-'a')], w)
	}
	for _, c := range s {
		words := hash[int(byte(c)-'a')]
		hash[int(byte(c)-'a')] = []string{}
		for _, w := range words {
			if len(w) == 1 {
				res += 1
				continue
			}
			hash[int(w[1]-'a')] = append(hash[int(w[1]-'a')], w[1:])
		}
	}
	return res
}
