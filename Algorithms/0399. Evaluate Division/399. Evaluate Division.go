package leetcode

type stringUnionFind struct {
	parents map[string]string
	vals    map[string]float64
}

func (suf stringUnionFind) add(x string) {
	if _, ok := suf.parents[x]; ok {
		return
	}
	suf.parents[x] = x
	suf.vals[x] = 1.0
}

func (suf stringUnionFind) find(x string) string {
	p := ""
	if v, ok := suf.parents[x]; ok {
		p = v
	} else {
		p = x
	}
	if x != p {
		pp := suf.find(p)
		suf.vals[x] *= suf.vals[p]
		suf.parents[x] = pp
	}
	if v, ok := suf.parents[x]; ok {
		return v
	}
	return x
}

func (suf stringUnionFind) union(x, y string, v float64) {
	suf.add(x)
	suf.add(y)
	px, py := suf.find(x), suf.find(y)
	suf.parents[px] = py
	// x / px = vals[x]
	// x / y = v
	// 由上面 2 个式子就可以得出 px = v * vals[y] / vals[x]
	suf.vals[px] = v * suf.vals[y] / suf.vals[x]
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res, suf := make([]float64, len(queries)), stringUnionFind{parents: map[string]string{}, vals: map[string]float64{}}
	for i := 0; i < len(values); i++ {
		suf.union(equations[i][0], equations[i][1], values[i])
	}
	for i := 0; i < len(queries); i++ {
		x, y := queries[i][0], queries[i][1]
		if _, ok := suf.parents[x]; ok {
			if _, ok := suf.parents[y]; ok {
				if suf.find(x) == suf.find(y) {
					res[i] = suf.vals[x] / suf.vals[y]
				} else {
					res[i] = -1
				}
			} else {
				res[i] = -1
			}
		} else {
			res[i] = -1
		}
	}
	return res
}
