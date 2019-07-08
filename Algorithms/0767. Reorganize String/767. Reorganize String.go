package leetcode

import (
	"sort"
)

func reorganizeString(S string) string {
	fs := frequencySort767(S)
	if fs == "" {
		return ""
	}
	bs := []byte(fs)
	ans := ""
	j := (len(bs)-1)/2 + 1
	for i := 0; i <= (len(bs)-1)/2; i++ {
		ans += string(bs[i])
		if j < len(bs) {
			ans += string(bs[j])
		}
		j++
	}
	return ans
}

func frequencySort767(s string) string {
	if s == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(s)
	for _, b := range sb {
		sMap[b]++
		if sMap[b] > (len(sb)+1)/2 {
			return ""
		}
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}

	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}
