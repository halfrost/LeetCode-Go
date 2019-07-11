package leetcode

func findTheDifference(s string, t string) byte {
	n, ch := len(t), t[len(t)-1]
	for i := 0; i < n-1; i++ {
		ch ^= s[i]
		ch ^= t[i]
	}
	return ch
}
