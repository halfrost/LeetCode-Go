package leetcode

func decode(encoded []int, first int) []int {
	arr := make([]int, len(encoded)+1)
	arr[0] = first
	for i, val := range encoded {
		arr[i+1] = arr[i] ^ val
	}
	return arr
}
