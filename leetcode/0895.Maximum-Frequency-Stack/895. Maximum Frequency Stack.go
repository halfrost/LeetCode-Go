package leetcode

type FreqStack struct {
	freq    map[int]int
	group   map[int][]int
	maxfreq int
}

func Constructor895() FreqStack {
	hash := make(map[int]int)
	maxHash := make(map[int][]int)
	return FreqStack{freq: hash, group: maxHash}
}

func (this *FreqStack) Push(x int) {
	if _, ok := this.freq[x]; ok {
		this.freq[x]++
	} else {
		this.freq[x] = 1
	}
	f := this.freq[x]
	if f > this.maxfreq {
		this.maxfreq = f
	}

	this.group[f] = append(this.group[f], x)
}

func (this *FreqStack) Pop() int {
	tmp := this.group[this.maxfreq]
	x := tmp[len(tmp)-1]
	this.group[this.maxfreq] = this.group[this.maxfreq][:len(this.group[this.maxfreq])-1]
	this.freq[x]--
	if len(this.group[this.maxfreq]) == 0 {
		this.maxfreq--
	}
	return x
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
