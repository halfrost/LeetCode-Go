package leetcode

// 解法一
func pushDominoes(dominoes string) string {
	d := []byte(dominoes)
	for i := 0; i < len(d); {
		j := i + 1
		for j < len(d)-1 && d[j] == '.' {
			j++
		}
		push(d[i : j+1])
		i = j
	}
	return string(d)
}

func push(d []byte) {
	first, last := 0, len(d)-1
	switch d[first] {
	case '.', 'L':
		if d[last] == 'L' {
			for ; first < last; first++ {
				d[first] = 'L'
			}
		}
	case 'R':
		if d[last] == '.' || d[last] == 'R' {
			for ; first <= last; first++ {
				d[first] = 'R'
			}
		} else if d[last] == 'L' {
			for first < last {
				d[first] = 'R'
				d[last] = 'L'
				first++
				last--
			}
		}
	}
}

// 解法二
func pushDominoes1(dominoes string) string {
	dominoes = "L" + dominoes + "R"
	res := ""
	for i, j := 0, 1; j < len(dominoes); j++ {
		if dominoes[j] == '.' {
			continue
		}
		if i > 0 {
			res += string(dominoes[i])
		}
		middle := j - i - 1
		if dominoes[i] == dominoes[j] {
			for k := 0; k < middle; k++ {
				res += string(dominoes[i])
			}
		} else if dominoes[i] == 'L' && dominoes[j] == 'R' {
			for k := 0; k < middle; k++ {
				res += string('.')
			}
		} else {
			for k := 0; k < middle/2; k++ {
				res += string('R')
			}
			for k := 0; k < middle%2; k++ {
				res += string('.')
			}
			for k := 0; k < middle/2; k++ {
				res += string('L')
			}
		}
		i = j
	}
	return res
}
