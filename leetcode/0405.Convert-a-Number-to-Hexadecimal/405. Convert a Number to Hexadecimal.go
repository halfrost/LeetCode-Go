package leetcode

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num += 1 << 32
	}
	mp := map[int]string{
		0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
		10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f",
	}
	var bitArr []string
	for num > 0 {
		bitArr = append(bitArr, mp[num%16])
		num /= 16
	}
	str := ""
	for i := len(bitArr) - 1; i >= 0; i-- {
		str += bitArr[i]
	}
	return str
}
