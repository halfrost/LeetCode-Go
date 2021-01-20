package leetcode

type OrderedStream struct {
	ptr    int
	stream []string
}

func Constructor(n int) OrderedStream {
	ptr, stream := 1, make([]string, n+1)
	return OrderedStream{ptr: ptr, stream: stream}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.stream[id] = value
	res := []string{}
	if this.ptr == id || this.stream[this.ptr] != "" {
		res = append(res, this.stream[this.ptr])
		for i := id + 1; i < len(this.stream); i++ {
			if this.stream[i] != "" {
				res = append(res, this.stream[i])
			} else {
				this.ptr = i
				return res
			}
		}
	}
	if len(res) > 0 {
		return res
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */
