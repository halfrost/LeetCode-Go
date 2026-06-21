package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem284(t *testing.T) {
	iter := &Iterator{}

	// cover the underlying Iterator methods directly
	if !iter.hasNext() {
		t.Fatalf("expected underlying iterator to have next")
	}
	if iter.next() != 0 {
		t.Fatalf("expected underlying iterator next to be 0")
	}

	pi := Constructor(iter)

	// hasNext when no element is cached -> delegates to iter.hasNext()
	if !pi.hasNext() {
		t.Fatalf("expected peeking iterator to have next")
	}
	fmt.Printf("hasNext (no cache) = %v\n", true)

	// next when no element is cached -> delegates to iter.next()
	if got := pi.next(); got != 0 {
		t.Fatalf("expected next to be 0, got %d", got)
	}
	fmt.Printf("next (no cache) = %d\n", 0)

	// peek when no element cached -> caches and returns
	if got := pi.peek(); got != 0 {
		t.Fatalf("expected peek to be 0, got %d", got)
	}
	fmt.Printf("peek (no cache) = %d\n", 0)

	// peek again -> returns cached element (hasEl branch)
	if got := pi.peek(); got != 0 {
		t.Fatalf("expected peek (cached) to be 0, got %d", got)
	}
	fmt.Printf("peek (cached) = %d\n", 0)

	// hasNext when element is cached (hasEl branch)
	if !pi.hasNext() {
		t.Fatalf("expected peeking iterator to have next when cached")
	}
	fmt.Printf("hasNext (cached) = %v\n", true)

	// next when element is cached -> returns cached and clears flag
	if got := pi.next(); got != 0 {
		t.Fatalf("expected next (cached) to be 0, got %d", got)
	}
	fmt.Printf("next (cached) = %d\n", 0)
}
