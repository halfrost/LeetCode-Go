package leetcode

import "fmt"

func largestTimeFromDigits(A []int) string {
	flag, res := false, 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				l := 6 - i - j - k
				hour := A[i]*10 + A[j]
				min := A[k]*10 + A[l]
				if hour < 24 && min < 60 {
					if hour*60+min >= res {
						res = hour*60 + min
						flag = true
					}
				}
			}
		}
	}
	if flag {
		return fmt.Sprintf("%02d:%02d", res/60, res%60)
	} else {
		return ""
	}
}
