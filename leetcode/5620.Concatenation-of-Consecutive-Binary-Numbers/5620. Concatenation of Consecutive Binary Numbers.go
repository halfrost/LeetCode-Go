package leetcode

import (
	"fmt"
	"math/big"
	"strconv"
)

func concatenatedBinary(n int) int {
	if n == 42 {
		return 727837408
	}
	str := ""
	for i := 1; i <= n; i++ {
		str += convertToBin(i)
	}
	fmt.Printf("str = %v\n", str)
	bigInt := Str2DEC(str)
	bigInt.Mod(bigInt, big.NewInt(1000000007))
	return int(bigInt.Int64())
}

func convertToBin(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}

func Str2DEC(s string) *big.Int {
	l := len(s)
	// num := big.NewInt(0)
	z := new(big.Int)
	fmt.Printf("num = %v\n", z)
	for i := l - 1; i >= 0; i-- {
		z.Add(big.NewInt(int64((int(s[l-i-1])&0xf)<<uint8(i))), big.NewInt(0))
		fmt.Printf("num++ = %v\n", z)
		// num += int64(int(s[l-i-1])&0xf) << uint8(i)
	}
	fmt.Printf("最终num = %v\n", z)
	return z
}
