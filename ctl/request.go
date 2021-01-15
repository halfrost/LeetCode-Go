package main

import (
	"bytes"
	"fmt"
	"github.com/mozillazg/request"
	"io/ioutil"
	"net/http"
)

const (
	// AllProblemURL define
	AllProblemURL = "https://leetcode.com/api/problems/all/"
	// QraphqlURL define
	QraphqlURL = "https://leetcode.com/graphql"
	// LoginPageURL define
	LoginPageURL = "https://leetcode.com/accounts/login/"
	// AlgorithmsURL define
	AlgorithmsURL = "https://leetcode.com/api/problems/Algorithms/"

	// ArrayProblemURL define
	ArrayProblemURL = "https://leetcode.com/tag/array/"
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
		"x-csrftoken":     cfg.CSRFtoken,
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
		return []byte{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("getRaw: Read Error: " + err.Error())
		return []byte{}
	}
	if resp.StatusCode == 200 {
		fmt.Println("Get problem Success!")
	}
	return body
}

func getProblemAllList() []byte {
	return getRaw(AllProblemURL)
}

// Variables define
type Variables struct {
	slug string
}

func getQraphql(payload string) []byte {
	req := newReq()
	resp, err := req.PostForm(QraphqlURL, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Printf("getRaw: Get Error: " + err.Error())
		return []byte{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("getRaw: Read Error: " + err.Error())
		return []byte{}
	}
	if resp.StatusCode == 200 {
		fmt.Println("Get problem Success!")
	}
	return body
}

func getTopicTag(variable string) string {
	return fmt.Sprintf(`{
		"operationName": "getTopicTag",
		"variables": {
			"slug": "%s"
		},
		"query": "query getTopicTag($slug: String!) {  topicTag(slug: $slug) {    name    translatedName    slug    questions {      status      questionId      questionFrontendId      title      titleSlug      translatedTitle      stats      difficulty      isPaidOnly      topicTags {        name        translatedName        slug        __typename      }      companyTags {        name        translatedName        slug        __typename      }      __typename    }    frequencies    __typename  }  favoritesLists {    publicFavorites {      ...favoriteFields      __typename    }    privateFavorites {      ...favoriteFields      __typename    }    __typename  }}fragment favoriteFields on FavoriteNode {  idHash  id  name  isPublicFavorite  viewCount  creator  isWatched  questions {    questionId    title    titleSlug    __typename  }  __typename}"
		}`, variable)
}

func getTagProblemList(tag string) []byte {
	return getQraphql(getTopicTag(tag))
}
