package main

import (
	"fmt"
	"github.com/mozillazg/request"
	"io/ioutil"
	"net/http"
)

const (
	// AllProblemURL define
	AllProblemURL = "https://leetcode.com/api/problems/all/"
	// LoginPageURL define
	LoginPageURL = "https://leetcode.com/accounts/login/"
	// AlgorithmsURL define
	AlgorithmsURL = "https://leetcode.com/api/problems/Algorithms/"
)

var req *request.Request

func newReq() *request.Request {
	if req == nil {
		req = signin()
	}
	return req
}

func signin() *request.Request {
	cfg := getConfig()
	req := request.NewRequest(new(http.Client))
	req.Headers = map[string]string{
		"Content-Type":    "application/json",
		"Accept-Encoding": "",
		"cookie":          cfg.Cookie,
		"Referer":         "https://leetcode.com/accounts/login/",
		"origin":          "https://leetcode.com",
	}
	return req
}

func getRaw(URL string) []byte {
	req := newReq()
	resp, err := req.Get(URL)
	if err != nil {
		fmt.Printf("getRaw: Get Error: " + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("getRaw: Read Error: " + err.Error())
	}
	return body
}

func getProblemAllList(URL string) []byte {
	req := newReq()
	resp, err := req.Get(URL)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	return body
}
