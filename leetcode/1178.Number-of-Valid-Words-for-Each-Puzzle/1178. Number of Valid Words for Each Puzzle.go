package leetcode

/*
	匹配跟单词中的字母顺序，字母个数都无关，可以用bitmap压缩
	1. 记录word中 利用map记录各种bit标示的个数
	2. puzzles 中各个字母都不相同! 记录bitmap，然后搜索子空间中各种bit标识的个数的和
	   因为puzzles长度最长是7，所以搜索空间 2^7
*/
func findNumOfValidWords(words []string, puzzles []string) []int {
	wordBitStatusMap, res := make(map[uint32]int, 0), []int{}
	for _, w := range words {
		wordBitStatusMap[toBitMap([]byte(w))]++
	}
	for _, p := range puzzles {
		var bitMap uint32
		var totalNum int
		bitMap |= (1 << (p[0] - 'a')) //work中要包含 p 的第一个字母 所以这个bit位上必须是1
		findNum([]byte(p)[1:], bitMap, &totalNum, wordBitStatusMap)
		res = append(res, totalNum)
	}
	return res
}

func toBitMap(word []byte) uint32 {
	var res uint32
	for _, b := range word {
		res |= (1 << (b - 'a'))
	}
	return res
}

//利用dfs 搜索 pussles的子空间
func findNum(puzzles []byte, bitMap uint32, totalNum *int, m map[uint32]int) {
	if len(puzzles) == 0 {
		*totalNum = *totalNum + m[bitMap]
		return
	}
	//不包含puzzles[0],即puzzles[0]对应bit是0
	findNum(puzzles[1:], bitMap, totalNum, m)
	//包含puzzles[0],即puzzles[0]对应bit是1
	bitMap |= (1 << (puzzles[0] - 'a'))
	findNum(puzzles[1:], bitMap, totalNum, m)
	bitMap ^= (1 << (puzzles[0] - 'a')) //异或 清零
	return
}
