package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 并查集优化搜索解法
func accountsMerge(accounts [][]string) (r [][]string) {
	uf := template.UnionFind{}
	uf.Init(len(accounts))
	// emailToID 将所有的 email 邮箱都拆开，拆开与 id(数组下标) 对应
	// idToName 将 id(数组下标) 与 name 对应
	// idToEmails 将 id(数组下标) 与整理好去重以后的 email 组对应
	emailToID, idToName, idToEmails, res := make(map[string]int), make(map[int]string), make(map[int][]string), [][]string{}
	for id, acc := range accounts {
		idToName[id] = acc[0]
		for i := 1; i < len(acc); i++ {
			pid, ok := emailToID[acc[i]]
			if ok {
				uf.Union(id, pid)
			}
			emailToID[acc[i]] = id
		}
	}
	for email, id := range emailToID {
		pid := uf.Find(id)
		idToEmails[pid] = append(idToEmails[pid], email)
	}
	for id, emails := range idToEmails {
		name := idToName[id]
		sort.Strings(emails)
		res = append(res, append([]string{name}, emails...))
	}
	return res
}

// 解法二 并查集暴力解法
func accountsMerge1(accounts [][]string) [][]string {
	if len(accounts) == 0 {
		return [][]string{}
	}
	uf, res, visited := template.UnionFind{}, [][]string{}, map[int]bool{}
	uf.Init(len(accounts))
	for i := 0; i < len(accounts); i++ {
		for j := i + 1; j < len(accounts); j++ {
			if accounts[i][0] == accounts[j][0] {
				tmpA, tmpB, flag := accounts[i][1:], accounts[j][1:], false
				for j := 0; j < len(tmpA); j++ {
					for k := 0; k < len(tmpB); k++ {
						if tmpA[j] == tmpB[k] {
							flag = true
							break
						}
					}
					if flag {
						break
					}
				}
				if flag {
					uf.Union(i, j)
				}
			}
		}
	}
	for i := 0; i < len(accounts); i++ {
		if visited[i] {
			continue
		}
		emails, account, tmpMap := accounts[i][1:], []string{accounts[i][0]}, map[string]string{}
		for j := i + 1; j < len(accounts); j++ {
			if uf.Find(j) == uf.Find(i) {
				visited[j] = true
				for _, v := range accounts[j][1:] {
					tmpMap[v] = v
				}
			}
		}
		for _, v := range emails {
			tmpMap[v] = v
		}
		emails = []string{}
		for key := range tmpMap {
			emails = append(emails, key)
		}
		sort.Strings(emails)
		account = append(account, emails...)
		res = append(res, account)
	}
	return res
}
