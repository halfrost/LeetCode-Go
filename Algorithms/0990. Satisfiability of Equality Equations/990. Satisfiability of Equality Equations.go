package leetcode

func equationsPossible(equations []string) bool {
	if len(equations) == 0 {
		return false
	}
	uf := UnionFind{}
	uf.init(26)
	for _, equ := range equations {
		if equ[1] == '=' && equ[2] == '=' {
			uf.union(int(equ[0]-'a'), int(equ[3]-'a'))
		}
	}
	for _, equ := range equations {
		if equ[1] == '!' && equ[2] == '=' {
			if uf.find(int(equ[0]-'a')) == uf.find(int(equ[3]-'a')) {
				return false
			}
		}
	}
	return true
}
