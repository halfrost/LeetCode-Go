package leetcode

import "sort"

type data struct {
	time  int
	value string
}

// TimeMap is a timebased key-value store
// TimeMap define
type TimeMap map[string][]data

// Constructor981 define
func Constructor981() TimeMap {
	return make(map[string][]data, 1024)
}

// Set define
func (t TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := t[key]; !ok {
		t[key] = make([]data, 1, 1024)
	}
	t[key] = append(t[key], data{
		time:  timestamp,
		value: value,
	})
}

// Get define
func (t TimeMap) Get(key string, timestamp int) string {
	d := t[key]
	i := sort.Search(len(d), func(i int) bool {
		return timestamp < d[i].time
	})
	i--
	return t[key][i].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
