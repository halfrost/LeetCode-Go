package leetcode

import (
	"sort"
)

// 解法一 二分搜索
func kthSmallestPrimeFraction(A []int, K int) []int {
	low, high, n := 0.0, 1.0, len(A)
	// 因为是在小数内使用二分查找，无法像在整数范围内那样通过 mid+1 和边界判断来终止循环
	// 所以此处根据 count 来结束循环
	for {
		mid, count, p, q, j := (high+low)/2.0, 0, 0, 1, 0
		for i := 0; i < n; i++ {
			for j < n && float64(A[i]) > float64(mid)*float64(A[j]) {
				j++
			}
			count += n - j
			if j < n && q*A[i] > p*A[j] {
				p = A[i]
				q = A[j]
			}
		}
		if count == K {
			return []int{p, q}
		} else if count < K {
			low = mid
		} else {
			high = mid
		}
	}
}

// 解法二 暴力解法，时间复杂度 O(n^2)
func kthSmallestPrimeFraction1(A []int, K int) []int {
	if len(A) == 0 || (len(A)*(len(A)-1))/2 < K {
		return []int{}
	}
	fractions := []Fraction{}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			fractions = append(fractions, Fraction{molecule: A[i], denominator: A[j]})
		}
	}
	sort.Sort(SortByFraction(fractions))
	return []int{fractions[K-1].molecule, fractions[K-1].denominator}
}

// Fraction define
type Fraction struct {
	molecule    int
	denominator int
}

// SortByFraction define
type SortByFraction []Fraction

func (a SortByFraction) Len() int      { return len(a) }
func (a SortByFraction) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByFraction) Less(i, j int) bool {
	return a[i].molecule*a[j].denominator < a[j].molecule*a[i].denominator
}
