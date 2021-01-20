# [1694. Reformat Phone Number](https://leetcode.com/problems/reformat-phone-number/)


## 题目

You are given a phone number as a string `number`. `number` consists of digits, spaces `' '`, and/or dashes `'-'`.

You would like to reformat the phone number in a certain manner. Firstly, **remove** all spaces and dashes. Then, **group** the digits from left to right into blocks of length 3 **until** there are 4 or fewer digits. The final digits are then grouped as follows:

- 2 digits: A single block of length 2.
- 3 digits: A single block of length 3.
- 4 digits: Two blocks of length 2 each.

The blocks are then joined by dashes. Notice that the reformatting process should **never** produce any blocks of length 1 and produce **at most** two blocks of length 2.

Return *the phone number after formatting.*

**Example 1:**

```
Input: number = "1-23-45 6"
Output: "123-456"
Explanation: The digits are "123456".
Step 1: There are more than 4 digits, so group the next 3 digits. The 1st block is "123".
Step 2: There are 3 digits remaining, so put them in a single block of length 3. The 2nd block is "456".
Joining the blocks gives "123-456".

```

**Example 2:**

```
Input: number = "123 4-567"
Output: "123-45-67"
Explanation: The digits are "1234567".
Step 1: There are more than 4 digits, so group the next 3 digits. The 1st block is "123".
Step 2: There are 4 digits left, so split them into two blocks of length 2. The blocks are "45" and "67".
Joining the blocks gives "123-45-67".

```

**Example 3:**

```
Input: number = "123 4-5678"
Output: "123-456-78"
Explanation: The digits are "12345678".
Step 1: The 1st block is "123".
Step 2: The 2nd block is "456".
Step 3: There are 2 digits left, so put them in a single block of length 2. The 3rd block is "78".
Joining the blocks gives "123-456-78".

```

**Example 4:**

```
Input: number = "12"
Output: "12"

```

**Example 5:**

```
Input: number = "--17-5 229 35-39475 "
Output: "175-229-353-94-75"

```

**Constraints:**

- `2 <= number.length <= 100`
- `number` consists of digits and the characters `'-'` and `' '`.
- There are at least **two** digits in `number`.

## 题目大意

给你一个字符串形式的电话号码 number 。number 由数字、空格 ' '、和破折号 '-' 组成。

请你按下述方式重新格式化电话号码。

- 首先，删除 所有的空格和破折号。
- 其次，将数组从左到右 每 3 个一组 分块，直到 剩下 4 个或更少数字。剩下的数字将按下述规定再分块：
    - 2 个数字：单个含 2 个数字的块。
    - 3 个数字：单个含 3 个数字的块。
    - 4 个数字：两个分别含 2 个数字的块。

最后用破折号将这些块连接起来。注意，重新格式化过程中 不应该 生成仅含 1 个数字的块，并且 最多 生成两个含 2 个数字的块。返回格式化后的电话号码。

## 解题思路

- 简单题。先判断号码是不是 2 位和 4 位，如果是，单独输出这 2 种情况。剩下的都是 3 位以上了，取余，判断剩余的数字是 2 个还是 4 个。这时不可能存在剩 1 位数的情况。除 3 余 1，即剩 4 位的情况，末尾 4 位需要 2 个一组输出。除 3 余 2，即剩  2 位的情况。处理好末尾，再逆序每 3 个一组连接 "-" 即可。可能需要注意的 case 见 test 文件。

## 代码

```go
package leetcode

import (
	"strings"
)

func reformatNumber(number string) string {
	parts, nums := []string{}, []rune{}
	for _, r := range number {
		if r != '-' && r != ' ' {
			nums = append(nums, r)
		}
	}
	threeDigits, twoDigits := len(nums)/3, 0
	switch len(nums) % 3 {
	case 1:
		threeDigits--
		twoDigits = 2
	case 2:
		twoDigits = 1
	default:
		twoDigits = 0
	}
	for i := 0; i < threeDigits; i++ {
		s := ""
		s += string(nums[0:3])
		nums = nums[3:]
		parts = append(parts, s)
	}
	for i := 0; i < twoDigits; i++ {
		s := ""
		s += string(nums[0:2])
		nums = nums[2:]
		parts = append(parts, s)
	}
	return strings.Join(parts, "-")
}
```