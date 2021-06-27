package leetcode

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	targetNum, visited := strToInt(target), make([]bool, 10000)
	visited[0] = true
	for _, deadend := range deadends {
		num := strToInt(deadend)
		if num == 0 {
			return -1
		}
		visited[num] = true
	}
	depth, curDepth, nextDepth := 0, []int16{0}, make([]int16, 0)
	var nextNum int16
	for len(curDepth) > 0 {
		nextDepth = nextDepth[0:0]
		for _, curNum := range curDepth {
			for incrementer := int16(1000); incrementer > 0; incrementer /= 10 {
				digit := (curNum / incrementer) % 10
				if digit == 9 {
					nextNum = curNum - 9*incrementer
				} else {
					nextNum = curNum + incrementer
				}
				if nextNum == targetNum {
					return depth + 1
				}
				if !visited[nextNum] {
					visited[nextNum] = true
					nextDepth = append(nextDepth, nextNum)
				}
				if digit == 0 {
					nextNum = curNum + 9*incrementer
				} else {
					nextNum = curNum - incrementer
				}
				if nextNum == targetNum {
					return depth + 1
				}
				if !visited[nextNum] {
					visited[nextNum] = true
					nextDepth = append(nextDepth, nextNum)
				}
			}
		}
		curDepth, nextDepth = nextDepth, curDepth
		depth++
	}
	return -1
}

func strToInt(str string) int16 {
	return int16(str[0]-'0')*1000 + int16(str[1]-'0')*100 + int16(str[2]-'0')*10 + int16(str[3]-'0')
}
