package leetcode

import (
	"fmt"
)

func oddEvenJumps(A []int) int {
	oddJumpMap, evenJumpMap, current, res := map[int]int{}, map[int]int{}, 0, 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if v, ok := oddJumpMap[i]; ok {
				if A[i] <= A[j] && A[j] <= A[v] {
					if A[j] == A[v] && j < oddJumpMap[i] {
						oddJumpMap[i] = j
					} else if A[j] < A[v] {
						oddJumpMap[i] = j
					}
				}
			} else {
				if A[i] <= A[j] {
					oddJumpMap[i] = j
				}
			}
		}
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if v, ok := evenJumpMap[i]; ok {
				if A[i] >= A[j] && A[j] >= A[v] {
					if A[j] == A[v] && j < evenJumpMap[i] {
						evenJumpMap[i] = j
					} else if A[j] > A[v] {
						evenJumpMap[i] = j
					}
				}
			} else {
				if A[i] >= A[j] {
					evenJumpMap[i] = j
				}
			}
		}
	}
	fmt.Printf("oddJumpMap = %v evenJumpMap = %v\n", oddJumpMap, evenJumpMap)
	for i := 0; i < len(A); i++ {
		count := 1
		current = i
		for {
			if count%2 == 1 {
				if v, ok := oddJumpMap[current]; ok {
					if v == len(A)-1 {
						res++
						break
					}
					current = v
					count++
				} else {
					break
				}
			}
			if count%2 == 0 {
				if v, ok := evenJumpMap[current]; ok {
					if v == len(A)-1 {
						res++
						break
					}
					current = v
					count++
				} else {
					break
				}
			}
		}
	}
	return res + 1
}
