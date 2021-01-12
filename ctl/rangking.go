package main

import (
	"fmt"
	"strconv"
	"strings"
)

// getRanking 让这个方法优雅一点
func getRanking() int {
	// 获取网页数据
	URL := fmt.Sprintf("https://leetcode.com/%s/", getConfig().Username)
	data := getRaw(URL)
	str := string(data)
	// 通过不断裁剪 str 获取排名信息
	fmt.Println(str)
	i := strings.Index(str, "ng-init")
	j := i + strings.Index(str[i:], "ng-cloak")
	str = str[i:j]
	i = strings.Index(str, "(")
	j = strings.Index(str, ")")
	str = str[i:j]
	//	fmt.Println("2\n", str)
	strs := strings.Split(str, ",")
	str = strs[6]
	//	fmt.Println("1\n", str)
	i = strings.Index(str, "'")
	j = 2 + strings.Index(str[2:], "'")
	//	fmt.Println("0\n", str)
	str = str[i+1 : j]
	r, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("无法把 %s 转换成数字Ranking", str)
	}
	return r
}
