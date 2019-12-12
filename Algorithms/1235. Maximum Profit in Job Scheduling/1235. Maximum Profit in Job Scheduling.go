package leetcode

import "sort"

type job struct {
	startTime int
	endTime   int
	profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs, dp := []job{}, make([]int, len(startTime))
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, job{startTime: startTime[i], endTime: endTime[i], profit: profit[i]})
	}
	sort.Sort(sortJobs(jobs))
	dp[0] = jobs[0].profit
	for i := 1; i < len(jobs); i++ {
		low, high := 0, i-1
		for low < high {
			mid := low + (high-low)>>1
			if jobs[mid+1].endTime <= jobs[i].startTime {
				low = mid + 1
			} else {
				high = mid
			}
		}
		if jobs[low].endTime <= jobs[i].startTime {
			dp[i] = max(dp[i-1], dp[low]+jobs[i].profit)
		} else {
			dp[i] = max(dp[i-1], jobs[i].profit)
		}
	}
	return dp[len(startTime)-1]
}

type sortJobs []job

func (s sortJobs) Len() int {
	return len(s)
}
func (s sortJobs) Less(i, j int) bool {
	if s[i].endTime == s[j].endTime {
		return s[i].profit < s[j].profit
	}
	return s[i].endTime < s[j].endTime
}
func (s sortJobs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
