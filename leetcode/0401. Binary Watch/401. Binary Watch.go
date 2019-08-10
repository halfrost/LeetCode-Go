package leetcode

import (
	"fmt"
	"strconv"
)

var (
	hour    = []string{"1", "2", "4", "8"}
	minute  = []string{"01", "02", "04", "08", "16", "32"}
	hourMap = map[int][]string{
		0: []string{"0"},
		1: []string{"1", "2", "4", "8"},
		2: []string{"3", "5", "9", "6", "10"},
		3: []string{"7", "11"},
	}
	minuteMap = map[int][]string{
		0: []string{"00"},
		1: []string{"01", "02", "04", "08", "16", "32"},
		2: []string{"03", "05", "09", "17", "33", "06", "10", "18", "34", "12", "20", "36", "24", "40", "48"},
		3: []string{"07", "11", "19", "35", "13", "21", "37", "25", "41", "49", "14", "22", "38", "26", "42", "50", "28", "44", "52", "56"},
		4: []string{"15", "23", "39", "27", "43", "51", "29", "45", "53", "57", "30", "46", "54", "58"},
		5: []string{"31", "47", "55", "59"},
	}
)

func readBinaryWatch(num int) []string {
	if num > 8 {
		return []string{}
	}
	res := []string{}
	for i := 0; i <= num; i++ {
		for j := 0; j < len(hourMap[i]); j++ {
			for k := 0; k < len(minuteMap[num-i]); k++ {
				res = append(res, hourMap[i][j]+":"+minuteMap[num-i][k])
			}
		}
	}
	return res
}

/// ---------------------------------------
/// ---------------------------------------
/// ---------------------------------------
/// ---------------------------------------
/// ---------------------------------------
// 以下是打表用到的函数
// 调用 findReadBinaryWatchMinute(num, 0, c, &res) 打表
func findReadBinaryWatchMinute(target, index int, c []int, res *[]string) {
	if target == 0 {
		str, tmp := "", 0
		for i := 0; i < len(c); i++ {
			t, _ := strconv.Atoi(minute[c[i]])
			tmp += t
		}
		if tmp < 10 {
			str = "0" + strconv.Itoa(tmp)
		} else {
			str = strconv.Itoa(tmp)
		}
		// fmt.Printf("找到解了 c = %v str = %v\n", c, str)
		fmt.Printf("\"%v\", ", str)
		return
	}
	for i := index; i < 6; i++ {
		c = append(c, i)
		findReadBinaryWatchMinute(target-1, i+1, c, res)
		c = c[:len(c)-1]
	}
}

// 调用 findReadBinaryWatchHour(num, 0, c, &res) 打表
func findReadBinaryWatchHour(target, index int, c []int, res *[]string) {
	if target == 0 {
		str, tmp := "", 0
		for i := 0; i < len(c); i++ {
			t, _ := strconv.Atoi(hour[c[i]])
			tmp += t
		}
		str = strconv.Itoa(tmp)
		//fmt.Printf("找到解了 c = %v str = %v\n", c, str)
		fmt.Printf("\"%v\", ", str)
		return
	}
	for i := index; i < 4; i++ {
		c = append(c, i)
		findReadBinaryWatchHour(target-1, i+1, c, res)
		c = c[:len(c)-1]
	}
}
