# [535. Encode and Decode TinyURL](https://leetcode.com/problems/encode-and-decode-tinyurl/)


## 题目

> Note: This is a companion problem to the System Design problem: Design TinyURL.

TinyURL is a URL shortening service where you enter a URL such as `https://leetcode.com/problems/design-tinyurl` and it returns a short URL such as `http://tinyurl.com/4e9iAk`.

Design the `encode` and `decode` methods for the TinyURL service. There is no restriction on how your encode/decode algorithm should work. You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.

## 题目大意

TinyURL是一种URL简化服务， 比如：当你输入一个URL [https://leetcode.com/problems/design-tinyurl](https://leetcode.com/problems/design-tinyurl) 时，它将返回一个简化的URL [http://tinyurl.com/4e9iAk](http://tinyurl.com/4e9iAk).

要求：设计一个 TinyURL 的加密 encode 和解密 decode 的方法。你的加密和解密算法如何设计和运作是没有限制的，你只需要保证一个URL可以被加密成一个TinyURL，并且这个TinyURL可以用解密方法恢复成原本的URL。

## 解题思路

- 简单题。由于题目并无规定 `encode()` 算法，所以自由度非常高。最简单的做法是把原始 `URL` 存起来，并记录下存在字符串数组中的下标位置。`decode()` 的时候根据存储的下标还原原始的 `URL`。

## 代码

```go
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
```