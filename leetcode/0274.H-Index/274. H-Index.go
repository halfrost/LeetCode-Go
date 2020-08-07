package leetcode

// 解法一
func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}
	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}
	return 0
}

// 解法二
func hIndex1(citations []int) int {
	quickSort164(citations, 0, len(citations)-1)
	hIndex := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= len(citations)-i {
			hIndex++
		} else {
			break
		}
	}
	return hIndex
}

func quickSort164(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition164(a, lo, hi)
	quickSort164(a, lo, p-1)
	quickSort164(a, p+1, hi)
}

func partition164(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
