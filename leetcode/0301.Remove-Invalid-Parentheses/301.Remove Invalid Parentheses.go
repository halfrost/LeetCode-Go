package leetcode

var (
	res      []string
	mp       map[string]int
	n        int
	length   int
	maxScore int
	str      string
)

func removeInvalidParentheses(s string) []string {
	lmoves, rmoves, lcnt, rcnt := 0, 0, 0, 0
	for _, v := range s {
		if v == '(' {
			lmoves++
			lcnt++
		} else if v == ')' {
			if lmoves != 0 {
				lmoves--
			} else {
				rmoves++
			}
			rcnt++
		}
	}
	n = len(s)
	length = n - lmoves - rmoves
	res = []string{}
	mp = make(map[string]int)
	maxScore = min(lcnt, rcnt)
	str = s
	backtrace(0, "", lmoves, rmoves, 0)
	return res
}

func backtrace(i int, cur string, lmoves int, rmoves int, score int) {
	if lmoves < 0 || rmoves < 0 || score < 0 || score > maxScore {
		return
	}
	if lmoves == 0 && rmoves == 0 {
		if len(cur) == length {
			if _, ok := mp[cur]; !ok {
				res = append(res, cur)
				mp[cur] = 1
			}
			return
		}
	}
	if i == n {
		return
	}
	if str[i] == '(' {
		backtrace(i+1, cur+string('('), lmoves, rmoves, score+1)
		backtrace(i+1, cur, lmoves-1, rmoves, score)
	} else if str[i] == ')' {
		backtrace(i+1, cur+string(')'), lmoves, rmoves, score-1)
		backtrace(i+1, cur, lmoves, rmoves-1, score)
	} else {
		backtrace(i+1, cur+string(str[i]), lmoves, rmoves, score)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
