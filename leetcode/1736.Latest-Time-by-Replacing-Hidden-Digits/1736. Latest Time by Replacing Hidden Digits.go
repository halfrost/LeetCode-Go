package leetcode

func maximumTime(time string) string {
	timeb := []byte(time)
	if timeb[3] == '?' {
		timeb[3] = '5'
	}
	if timeb[4] == '?' {
		timeb[4] = '9'
	}
	if timeb[0] == '?' {
		if int(timeb[1]-'0') > 3 && int(timeb[1]-'0') < 10 {
			timeb[0] = '1'
		} else {
			timeb[0] = '2'
		}
	}
	if timeb[1] == '?' {
		timeb[1] = '9'
	}
	if timeb[0] == '2' && timeb[1] == '9' {
		timeb[1] = '3'
	}
	return string(timeb)
}
