package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	lenth := len(flowerbed)
	for i := 0; i < lenth && n > 0; i += 2 {
		if flowerbed[i] == 0 {
			if i+1 == lenth || flowerbed[i+1] == 0 {
				n--
			} else {
				i++
			}
		}
	}
	if n == 0 {
		return true
	}
	return false
}
