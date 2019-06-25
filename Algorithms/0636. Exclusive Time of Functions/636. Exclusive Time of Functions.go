package leetcode

import (
	"strconv"
	"strings"
)

type log struct {
	id    int
	order string
	time  int
}

func exclusiveTime(n int, logs []string) []int {
	res, lastLog, stack := make([]int, n), log{id: -1, order: "", time: 0}, []log{}
	for i := 0; i < len(logs); i++ {
		a := strings.Split(logs[i], ":")
		id, _ := strconv.Atoi(a[0])
		time, _ := strconv.Atoi(a[2])

		if (lastLog.order == "start" && a[1] == "start") || (lastLog.order == "start" && a[1] == "end") {
			res[lastLog.id] += time - lastLog.time
			if a[1] == "end" {
				res[lastLog.id]++
			}
		}
		if lastLog.order == "end" && a[1] == "end" {
			res[id] += time - lastLog.time
		}
		if lastLog.order == "end" && a[1] == "start" && len(stack) != 0 {
			res[stack[len(stack)-1].id] += time - lastLog.time - 1
		}
		if a[1] == "start" {
			stack = append(stack, log{id: id, order: a[1], time: time})
		} else {
			stack = stack[:len(stack)-1]
		}
		lastLog = log{id: id, order: a[1], time: time}
	}
	return res
}
