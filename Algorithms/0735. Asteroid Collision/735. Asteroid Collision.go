package leetcode

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	for _, v := range asteroids {
		for len(res) != 0 && res[len(res)-1] > 0 && res[len(res)-1] < -v {
			res = res[:len(res)-1]
		}
		if len(res) == 0 || v > 0 || res[len(res)-1] < 0 {
			res = append(res, v)
		} else if v < 0 && res[len(res)-1] == -v {
			res = res[:len(res)-1]
		}
	}
	return res
}
