package leetcode

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	bfs := barcodesFrequencySort(barcodes)
	if len(bfs) == 0 {
		return []int{}
	}
	res := []int{}
	j := (len(bfs)-1)/2 + 1
	for i := 0; i <= (len(bfs)-1)/2; i++ {
		res = append(res, bfs[i])
		if j < len(bfs) {
			res = append(res, bfs[j])
		}
		j++
	}
	return res
}

func barcodesFrequencySort(s []int) []int {
	if len(s) == 0 {
		return []int{}
	}
	sMap := map[int]int{}   // 统计每个数字出现的频次
	cMap := map[int][]int{} // 按照频次作为 key 排序
	for _, b := range s {
		sMap[b]++
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}
	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]int, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return res
}
