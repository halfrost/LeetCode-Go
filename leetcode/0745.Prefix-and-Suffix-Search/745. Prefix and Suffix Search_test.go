package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem745(t *testing.T) {
	obj := Constructor745([]string{"apple"})
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.F("a", "e")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param2 := obj.F("b", "")
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)

	obj2 := Constructor_745_([]string{"apple"})
	fmt.Printf("obj2 = %v\n", obj2)
	if got := obj2.F_("a", "e"); got != param1 {
		t.Fatalf("F_(\"a\", \"e\") = %v, want %v", got, param1)
	}
	if got := obj2.F_("b", ""); got != param2 {
		t.Fatalf("F_(\"b\", \"\") = %v, want %v", got, param2)
	}
}
