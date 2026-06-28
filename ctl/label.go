package main

import (
	"fmt"

	"github.com/halfrost/LeetCode-Go/ctl/util"
)

// 上/下页导航已改由 Hugo 主题在渲染时动态生成
// （themes/book/layouts/partials/docs/inject/content-after.html），
// 不再把导航 HTML 写进 markdown，因此原先 add-pre-next / del-pre-next
// 这一整套读写 markdown 的逻辑已全部移除。

var (
	chapterOneFileOrder = []string{"_index", "Data_Structure", "Algorithm", "Time_Complexity"}
	chapterOneMenuOrder = []string{"_index", "#关于作者", "Data_Structure", "Algorithm", "Time_Complexity"}
	chapterTwoFileOrder = []string{"_index", "Array", "String", "Two_Pointers", "Linked_List", "Stack", "Tree", "Dynamic_Programming", "Backtracking", "Depth_First_Search", "Breadth_First_Search",
		"Binary_Search", "Math", "Hash_Table", "Sorting", "Bit_Manipulation", "Union_Find", "Sliding_Window", "Segment_Tree", "Binary_Indexed_Tree"}
	chapterThreeFileOrder = []string{"_index", "Segment_Tree", "UnionFind", "LRUCache", "LFUCache"}

	chapterMap = map[string]map[string]string{
		"ChapterOne": {
			"_index":          "第一章 序章",
			"Data_Structure":  "1.1 数据结构知识",
			"Algorithm":       "1.2 算法知识",
			"Time_Complexity": "1.3 时间复杂度",
		},
		"ChapterTwo": {
			"_index":               "第二章 算法专题",
			"Array":                "2.01 Array",
			"String":               "2.02 String",
			"Two_Pointers":         "2.03 ✅ Two Pointers",
			"Linked_List":          "2.04 ✅ Linked List",
			"Stack":                "2.05 ✅ Stack",
			"Tree":                 "2.06 Tree",
			"Dynamic_Programming":  "2.07 Dynamic Programming",
			"Backtracking":         "2.08 ✅ Backtracking",
			"Depth_First_Search":   "2.09 Depth First Search",
			"Breadth_First_Search": "2.10 Breadth First Search",
			"Binary_Search":        "2.11 Binary Search",
			"Math":                 "2.12 Math",
			"Hash_Table":           "2.13 Hash Table",
			"Sorting":              "2.14 ✅ Sorting",
			"Bit_Manipulation":     "2.15 ✅ Bit Manipulation",
			"Union_Find":           "2.16 ✅ Union Find",
			"Sliding_Window":       "2.17 ✅ Sliding Window",
			"Segment_Tree":         "2.18 ✅ Segment Tree",
			"Binary_Indexed_Tree":  "2.19 ✅ Binary Indexed Tree",
		},
		"ChapterThree": {
			"_index":       "第三章 一些模板",
			"Segment_Tree": "3.1 Segment Tree",
			"UnionFind":    "3.2 UnionFind",
			"LRUCache":     "3.3 LRUCache",
			"LFUCache":     "3.4 LFUCache",
		},
		"ChapterFour": {
			"_index": "第四章 Leetcode 题解",
		},
	}
)

func getChapterFourFileOrder() ([]string, []int) {
	solutions, solutionIds := util.LoadChapterFourDir()
	chapterFourFileOrder := []string{"_index"}
	chapterFourFileOrder = append(chapterFourFileOrder, solutions...)
	fmt.Printf("ChapterFour 中包括 _index 有 %v 个文件, len(id) = %v\n", len(chapterFourFileOrder), len(solutionIds))
	return chapterFourFileOrder, solutionIds
}
