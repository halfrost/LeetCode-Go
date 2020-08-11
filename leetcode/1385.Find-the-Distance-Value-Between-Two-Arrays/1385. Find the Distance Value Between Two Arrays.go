package leetcode

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i := range arr1 {
		for j := range arr2 {
			if abs(arr1[i]-arr2[j]) <= d {
				break
			}
			if j == len(arr2)-1 {
				res++
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
