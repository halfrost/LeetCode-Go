package leetcode

import (
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	freqWord1, freq1, freqList1, freqWord2, freq2, freqList2, flag := map[byte]int{}, []int{}, map[int][]byte{}, map[byte]int{}, []int{}, map[int][]byte{}, false
	for i := 0; i < len(word1); i++ {
		freqWord1[word1[i]]++
	}
	for i := 0; i < len(word2); i++ {
		freqWord2[word2[i]]++
	}
	freqTemp1 := map[int]int{}
	for k, v := range freqWord1 {
		freqTemp1[v]++
		if list, ok := freqList1[v]; ok {
			list = append(list, k)
			freqList1[v] = list
		} else {
			list := []byte{}
			list = append(list, k)
			freqList1[v] = list
		}
	}
	for _, v := range freqTemp1 {
		freq1 = append(freq1, v)
	}
	freqTemp2 := map[int]int{}
	for k, v := range freqWord2 {
		freqTemp2[v]++
		if list, ok := freqList2[v]; ok {
			list = append(list, k)
			freqList2[v] = list
		} else {
			list := []byte{}
			list = append(list, k)
			freqList2[v] = list
		}
	}
	for _, v := range freqTemp2 {
		freq2 = append(freq2, v)
	}
	if len(freq1) != len(freq2) {
		return false
	}
	sort.Ints(freq1)
	sort.Ints(freq2)
	for i := 0; i < len(freq1); i++ {
		if freq1[i] != freq2[i] {
			flag = true
			break
		}
	}
	if flag == true {
		return false
	}
	flag = false
	// 频次相同，再判断字母交换是否合法存在
	for k, v := range freqWord1 {
		if list, ok := freqList2[v]; ok {
			for i := 0; i < len(list); i++ {
				if list[i] != k && list[i] != '0' {
					// 交换的字母不存在
					if _, ok := freqWord1[list[i]]; !ok {
						flag = true
						break
					} else {
						// 交换的字母存在，重置这一位，代表这一个字母被交换了，下次不用它
						list[i] = '0'
					}
				}
			}
		} else {
			// 出现频次个数相同，但是频次不同
			flag = true
			break
		}
	}
	if flag == true {
		return false
	}
	return true
}
