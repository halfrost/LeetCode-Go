package leetcode

// 解法一 打表
func countVowelStrings(n int) int {
	res := []int{1, 5, 15, 35, 70, 126, 210, 330, 495, 715, 1001, 1365, 1820, 2380, 3060, 3876, 4845, 5985, 7315, 8855, 10626, 12650, 14950, 17550, 20475, 23751, 27405, 31465, 35960, 40920, 46376, 52360, 58905, 66045, 73815, 82251, 91390, 101270, 111930, 123410, 135751, 148995, 163185, 178365, 194580, 211876, 230300, 249900, 270725, 292825, 316251}
	return res[n]
}

func makeTable() []int {
	res, array := 0, []int{}
	for i := 0; i < 51; i++ {
		countVowelStringsDFS(i, 0, []string{}, []string{"a", "e", "i", "o", "u"}, &res)
		array = append(array, res)
		res = 0
	}
	return array
}

func countVowelStringsDFS(n, index int, cur []string, vowels []string, res *int) {
	vowels = vowels[index:]
	if len(cur) == n {
		(*res)++
		return
	}
	for i := 0; i < len(vowels); i++ {
		cur = append(cur, vowels[i])
		countVowelStringsDFS(n, i, cur, vowels, res)
		cur = cur[:len(cur)-1]
	}
}

// 解法二 数学方法 —— 组合
func countVowelStrings1(n int) int {
	return (n + 1) * (n + 2) * (n + 3) * (n + 4) / 24
}
