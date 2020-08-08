# [753. Cracking the Safe](https://leetcode.com/problems/cracking-the-safe/)



## 题目

There is a box protected by a password. The password is a sequence of `n` digits where each digit can be one of the first `k` digits `0, 1, ..., k-1`.

While entering a password, the last `n` digits entered will automatically be matched against the correct password.

For example, assuming the correct password is `"345"`, if you type `"012345"`, the box will open because the correct password matches the suffix of the entered password.

Return any password of **minimum length** that is guaranteed to open the box at some point of entering it.

**Example 1**:

```
Input: n = 1, k = 2
Output: "01"
Note: "10" will be accepted too.
```

**Example 2**:

```
Input: n = 2, k = 2
Output: "00110"
Note: "01100", "10011", "11001" will be accepted too.
```

**Note**:

1. `n` will be in the range `[1, 4]`.
2. `k` will be in the range `[1, 10]`.
3. `k^n` will be at most `4096`.


## 题目大意

有一个需要密码才能打开的保险箱。密码是 n 位数, 密码的每一位是 k 位序列 0, 1, ..., k-1 中的一个 。你可以随意输入密码，保险箱会自动记住最后 n 位输入，如果匹配，则能够打开保险箱。举个例子，假设密码是 "345"，你可以输入 "012345" 来打开它，只是你输入了 6 个字符.请返回一个能打开保险箱的最短字符串。

提示：

- n 的范围是 [1, 4]。
- k 的范围是 [1, 10]。
- k^n 最大可能为 4096。


## 解题思路

- 给出 2 个数字 n 和 k，n 代表密码是 n 位数，k 代表密码是 k 位。保险箱会记住最后 n 位输入。返回一个能打开保险箱的最短字符串。
- 看到题目中的数据范围，数据范围很小，所以可以考虑用 DFS。想解开保险箱，当然是暴力破解，枚举所有可能。题目要求我们输出一个最短的字符串，这里是本题的关键，为何有最短呢？这里有贪心的思想。如果下一次递归可以利用上一次的 n-1 位，那么最终输出的字符串肯定是最短的。(笔者这里就不证明了)，例如，例子 2 中，最短的字符串是 00，01，11，10。每次尝试都利用前一次的 n-1 位。想通了这个问题，利用 DFS 暴力回溯即可。

## 代码

```go
const number = "0123456789"

func crackSafe(n int, k int) string {
	if n == 1 {
		return number[:k]
	}
	visit, total := map[string]bool{}, int(math.Pow(float64(k), float64(n)))
	str := make([]byte, 0, total+n-1)
	for i := 1; i != n; i++ {
		str = append(str, '0')
	}
	dfsCrackSafe(total, n, k, &str, &visit)
	return string(str)
}

func dfsCrackSafe(depth, n, k int, str *[]byte, visit *map[string]bool) bool {
	if depth == 0 {
		return true
	}
	for i := 0; i != k; i++ {
		*str = append(*str, byte('0'+i))
		cur := string((*str)[len(*str)-n:])
		if _, ok := (*visit)[cur]; ok != true {
			(*visit)[cur] = true
			if dfsCrackSafe(depth-1, n, k, str, visit) {
				// 只有这里不需要删除
				return true
			}
			delete(*visit, cur)
		}
		// 删除
		*str = (*str)[0 : len(*str)-1]
	}
	return false
}
```