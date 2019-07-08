package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

type atom struct {
	name string
	cnt  int
}

type atoms []atom

func (this atoms) Len() int           { return len(this) }
func (this atoms) Less(i, j int) bool { return strings.Compare(this[i].name, this[j].name) < 0 }
func (this atoms) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this atoms) String() string {
	s := ""
	for _, a := range this {
		s += a.name
		if a.cnt > 1 {
			s += strconv.Itoa(a.cnt)
		}
	}
	return s
}

func countOfAtoms(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	stack := make([]string, 0)
	for i := 0; i < n; i++ {
		c := s[i]
		if c == '(' || c == ')' {
			stack = append(stack, string(c))
		} else if isUpperLetter(c) {
			j := i + 1
			for ; j < n; j++ {
				if !isLowerLetter(s[j]) {
					break
				}
			}
			stack = append(stack, s[i:j])
			i = j - 1
		} else if isDigital(c) {
			j := i + 1
			for ; j < n; j++ {
				if !isDigital(s[j]) {
					break
				}
			}
			stack = append(stack, s[i:j])
			i = j - 1
		}
	}

	cnt, deep := make([]map[string]int, 100), 0
	for i := 0; i < 100; i++ {
		cnt[i] = make(map[string]int)
	}
	for i := 0; i < len(stack); i++ {
		t := stack[i]
		if isUpperLetter(t[0]) {
			num := 1
			if i+1 < len(stack) && isDigital(stack[i+1][0]) {
				num, _ = strconv.Atoi(stack[i+1])
				i++
			}
			cnt[deep][t] += num
		} else if t == "(" {
			deep++
		} else if t == ")" {
			num := 1
			if i+1 < len(stack) && isDigital(stack[i+1][0]) {
				num, _ = strconv.Atoi(stack[i+1])
				i++
			}
			for k, v := range cnt[deep] {
				cnt[deep-1][k] += v * num
			}
			cnt[deep] = make(map[string]int)
			deep--
		}
	}
	as := atoms{}
	for k, v := range cnt[0] {
		as = append(as, atom{name: k, cnt: v})
	}
	sort.Sort(as)
	return as.String()
}

func isDigital(v byte) bool {
	if v >= '0' && v <= '9' {
		return true
	}
	return false
}

func isUpperLetter(v byte) bool {
	if v >= 'A' && v <= 'Z' {
		return true
	}
	return false
}

func isLowerLetter(v byte) bool {
	if v >= 'a' && v <= 'z' {
		return true
	}
	return false
}
