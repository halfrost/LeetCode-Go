package leetcode

func winnerOfGame(colors string) bool {
	As, Bs := 0, 0
	Acont, Bcont := 0, 0
	for _, color := range colors {
		if color == 'A' {
			Acont += 1
			Bcont = 0
		} else {
			Bcont += 1
			Acont = 0
		}
		if Acont >= 3 {
			As += Acont - 2
		}
		if Bcont >= 3 {
			Bs += Bcont - 2
		}
	}
	if As > Bs {
		return true
	}
	return false
}
