package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	urls []string
}

func Constructor() Codec {
	return Codec{[]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.urls = append(this.urls, longUrl)
	return "http://tinyurl.com/" + fmt.Sprintf("%v", len(this.urls)-1)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	tmp := strings.Split(shortUrl, "/")
	i, _ := strconv.Atoi(tmp[len(tmp)-1])
	return this.urls[i]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
