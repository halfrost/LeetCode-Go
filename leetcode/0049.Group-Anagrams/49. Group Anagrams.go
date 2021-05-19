package leetcode

import "sort"

// 方法1 桶排序
func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, s := range str {
			cnt[s-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	res := make([][]string, 0, len(mp))
	for _, r := range mp {
		res = append(res, r)
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// 方法2 排序
func groupAnagramsI(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRunes(sByte))
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}
