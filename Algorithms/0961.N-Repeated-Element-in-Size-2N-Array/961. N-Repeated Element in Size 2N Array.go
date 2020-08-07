package leetcode

func repeatedNTimes(A []int) int {
	kv := make(map[int]struct{})
	for _, val := range A {
		if _, ok := kv[val]; ok {
			return val
		}
		kv[val] = struct{}{}
	}
	return 0
}
