package leetcode

// 解法一
func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

// 解法二
func isAnagram1(s string, t string) bool {
	if s == "" && t == "" {
		return true
	}
	if s == "" || t == "" {
		return false
	}
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	quickSortByte(sBytes, 0, len(sBytes)-1)
	quickSortByte(tBytes, 0, len(tBytes)-1)

	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] != tBytes[i] {
			return false
		}
	}
	return true
}
func partitionByte(a []byte, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] > pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSortByte(a []byte, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionByte(a, lo, hi)
	quickSortByte(a, lo, p-1)
	quickSortByte(a, p+1, hi)
}
