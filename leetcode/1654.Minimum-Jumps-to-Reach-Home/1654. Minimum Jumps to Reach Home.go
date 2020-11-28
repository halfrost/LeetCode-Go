package leetcode

func minimumJumps(forbidden []int, a int, b int, x int) int {
	visited := make([]bool, 6000)
	for i := range forbidden {
		visited[forbidden[i]] = true
	}
	queue, res := [][2]int{{0, 0}}, -1
	for len(queue) > 0 {
		length := len(queue)
		res++
		for i := 0; i < length; i++ {
			cur, isBack := queue[i][0], queue[i][1]
			if cur == x {
				return res
			}
			if isBack == 0 && cur-b > 0 && !visited[cur-b] {
				visited[cur-b] = true
				queue = append(queue, [2]int{cur - b, 1})
			}
			if cur+a < len(visited) && !visited[cur+a] {
				visited[cur+a] = true
				queue = append(queue, [2]int{cur + a, 0})
			}
		}
		queue = queue[length:]
	}
	return -1
}
