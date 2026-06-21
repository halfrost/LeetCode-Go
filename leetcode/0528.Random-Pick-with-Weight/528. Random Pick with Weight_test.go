package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Problem528(t *testing.T) {
	w := []int{1, 3}
	sol := Constructor528(w)
	fmt.Printf("1.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("2.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("3.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("4.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("5.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("6.PickIndex = %v\n", sol.PickIndex())

	// Use weights with more elements and many iterations so the binary
	// search exercises every branch (low=mid+1, high=mid, and the
	// prefixSum[mid] == n early return).
	w2 := []int{3, 1, 1, 5, 2}
	sol2 := Constructor528(w2)
	total := w2[0] + w2[1] + w2[2] + w2[3] + w2[4]
	rand.Seed(1)
	for i := 0; i < 2000; i++ {
		idx := sol2.PickIndex()
		if idx < 0 || idx >= len(w2) {
			t.Fatalf("PickIndex returned out-of-range index %d for total %d", idx, total)
		}
	}
}
