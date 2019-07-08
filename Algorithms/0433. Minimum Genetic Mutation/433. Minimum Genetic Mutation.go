package leetcode

// 解法一 BFS
func minMutation(start string, end string, bank []string) int {
	wordMap, que, depth := getWordMap(bank, start), []string{start}, 0
	for len(que) > 0 {
		depth++
		qlen := len(que)
		for i := 0; i < qlen; i++ {
			word := que[0]
			que = que[1:]
			candidates := getCandidates433(word)
			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == end {
						return depth
					}
					delete(wordMap, candidate)
					que = append(que, candidate)
				}
			}
		}
	}
	return -1
}

func getCandidates433(word string) []string {
	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('A')+i) {
				res = append(res, word[:j]+string(int('A')+i)+word[j+1:])
			}
		}
	}
	return res
}

// 解法二 DFS
func minMutation1(start string, end string, bank []string) int {
	endGene := convert(end)
	startGene := convert(start)
	m := make(map[uint32]struct{})
	for _, gene := range bank {
		m[convert(gene)] = struct{}{}
	}
	if _, ok := m[endGene]; !ok {
		return -1
	}
	if check(startGene ^ endGene) {
		return 1
	}
	delete(m, startGene)
	step := make(map[uint32]int)
	step[endGene] = 0
	return dfsMutation(startGene, m, step)
}

func dfsMutation(start uint32, m map[uint32]struct{}, step map[uint32]int) int {
	if v, ok := step[start]; ok {
		return v
	}
	c := -1
	step[start] = c
	for k := range m {
		if check(k ^ start) {
			next := dfsMutation(k, m, step)
			if next != -1 {
				if c == -1 || c > next {
					c = next + 1
				}
			}
		}
	}
	step[start] = c
	return c
}

func check(val uint32) bool {
	if val == 0 {
		return false
	}
	if val&(val-1) == 0 {
		return true
	}
	for val > 0 {
		if val == 3 {
			return true
		}
		if val&3 != 0 {
			return false
		}
		val >>= 2
	}
	return false
}

func convert(gene string) uint32 {
	var v uint32
	for _, c := range gene {
		v <<= 2
		switch c {
		case 'C':
			v |= 1
		case 'G':
			v |= 2
		case 'T':
			v |= 3
		}
	}
	return v
}
