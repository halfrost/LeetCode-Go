package leetcode

import (
	"sort"
)

// TopVotedCandidate define
type TopVotedCandidate struct {
	persons []int
	times   []int
}

// Constructor911 define
func Constructor911(persons []int, times []int) TopVotedCandidate {
	leaders, votes := make([]int, len(persons)), make([]int, len(persons))
	leader := persons[0]
	for i := 0; i < len(persons); i++ {
		p := persons[i]
		votes[p]++
		if votes[p] >= votes[leader] {
			leader = p
		}
		leaders[i] = leader
	}
	return TopVotedCandidate{persons: leaders, times: times}
}

// Q define
func (tvc *TopVotedCandidate) Q(t int) int {
	i := sort.Search(len(tvc.times), func(p int) bool { return tvc.times[p] > t })
	return tvc.persons[i-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
