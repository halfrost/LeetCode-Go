package leetcode
/*
 * Author: Sk Shahriar Ahmed Raka
 * Email: skshahriarahmedraka@gmail.com
 * Telegram: https://t.me/shahriarraka
 * Github: https://github.com/skshahriarahmedraka
 * StackOverflow: https://stackoverflow.com/users/12216779/
 * Linkedin: https://linkedin.com/in/shahriarraka
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Copyright (c) 2022 Your Company
 * -----
 * HISTORY:
 */



// 58 / 58 test cases passed.
// Status: Accepted
// Runtime: 42 ms
// Memory Usage: 9.8 MB

// Submitted: 0 minutes ago
// Runtime: 42 ms, faster than 87.50% of Go online submissions for Minimum Number of Swaps to Make the String Balanced.
// Memory Usage: 9.8 MB, less than 81.25% of Go online submissions for Minimum Number of Swaps to Make the String Balanced.





func minSwaps(s string) int {
	maxClose:=0
	close:=0
	for _,i:= range s {
		if i=='['{
			close-=1
		}else {
			close+=1
		}
		maxClose=MAX(maxClose,close)
		

	}
	return (maxClose+1)/2

}

func MAX(i,j int)int{
	if i>j{
		return i 
	}
	return j
}