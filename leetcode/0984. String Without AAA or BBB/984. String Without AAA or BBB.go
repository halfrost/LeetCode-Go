package leetcode

func strWithout3a3b(A int, B int) string {
	ans, a, b := "", "a", "b"
	if B < A {
		A, B = B, A
		a, b = b, a
	}
	Dif := B - A
	if A == 1 && B == 1 { // ba
		ans = b + a
	} else if A == 1 && B < 5 { // bbabb
		for i := 0; i < B-2; i++ {
			ans = ans + b
		}
		ans = b + b + a + ans
	} else if (B-A)/A >= 1 { //bbabbabb
		for i := 0; i < A; i++ {
			ans = ans + b + b + a
			B = B - 2
		}
		for i := 0; i < B; i++ {
			ans = ans + b
		}
	} else { //bbabbabbabbabababa
		for i := 0; i < Dif; i++ {
			ans = ans + b + b + a
			B -= 2
			A--
		}
		for i := 0; i < B; i++ {
			ans = ans + b + a
		}
	}
	return ans
}
