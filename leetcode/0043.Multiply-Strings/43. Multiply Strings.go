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
 
 // 311 / 311 test cases passed.
 // Status: Accepted
 // Runtime: 0 ms
 // Memory Usage: 2.3 MB
 
 // You are here!
 // Your runtime beats 100.00 % of golang submissions.
 
 // You are here!
 // Your memory usage beats 77.90 % of golang submissions.
package leetcode

import (
	"math/big"

)


func multiply(num1 string, num2 string) string {
	var x big.Int
	var y big.Int
	var z big.Int 
	x.SetString(num1,10)
	y.SetString(num2,10)
	z.Mul(&x,&y)
	return z.String()
}
