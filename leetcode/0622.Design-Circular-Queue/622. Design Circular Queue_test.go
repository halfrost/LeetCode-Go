package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem622(t *testing.T) {
	obj := Constructor(3)
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.EnQueue(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param2 := obj.DeQueue()
	fmt.Printf("param_1 = %v obj = %v\n", param2, obj)
	param3 := obj.Front()
	fmt.Printf("param_1 = %v obj = %v\n", param3, obj)
	param4 := obj.Rear()
	fmt.Printf("param_1 = %v obj = %v\n", param4, obj)
	param5 := obj.IsEmpty()
	fmt.Printf("param_1 = %v obj = %v\n", param5, obj)
	param6 := obj.IsFull()
	fmt.Printf("param_1 = %v obj = %v\n", param6, obj)

	q := Constructor(3)
	// DeQueue on empty queue -> false
	if q.DeQueue() {
		t.Fatalf("DeQueue on empty should be false")
	}
	// Front/Rear on empty -> -1
	if q.Front() != -1 {
		t.Fatalf("Front on empty should be -1")
	}
	if q.Rear() != -1 {
		t.Fatalf("Rear on empty should be -1")
	}
	// Fill the queue
	if !q.EnQueue(10) {
		t.Fatalf("EnQueue 10 should be true")
	}
	if !q.EnQueue(20) {
		t.Fatalf("EnQueue 20 should be true")
	}
	if !q.EnQueue(30) {
		t.Fatalf("EnQueue 30 should be true")
	}
	// EnQueue on full -> false
	if q.EnQueue(40) {
		t.Fatalf("EnQueue on full should be false")
	}
	// Front returns first element
	if q.Front() != 10 {
		t.Fatalf("Front should be 10, got %v", q.Front())
	}
	// Rear returns last element (right-1 branch)
	if q.Rear() != 30 {
		t.Fatalf("Rear should be 30, got %v", q.Rear())
	}
	// DeQueue and wrap around so right becomes 0 to hit Rear's right==0 branch
	if !q.DeQueue() {
		t.Fatalf("DeQueue should be true")
	}
	if !q.EnQueue(40) {
		t.Fatalf("EnQueue 40 should be true after dequeue")
	}
	// now right wrapped to 0; Rear should be 40
	if q.Rear() != 40 {
		t.Fatalf("Rear should be 40, got %v", q.Rear())
	}
	fmt.Printf("q = %v\n", q)
}
