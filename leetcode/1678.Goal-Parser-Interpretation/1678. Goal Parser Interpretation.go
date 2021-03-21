package leetcode

func interpret(command string) string {
	if command == "" {
		return ""
	}
	res := ""
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res += "G"
		} else {
			if command[i] == '(' && command[i+1] == 'a' {
				res += "al"
				i += 3
			} else {
				res += "o"
				i++
			}
		}
	}
	return res
}
