package leetcode

import (
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	searchWordBytes, result := []byte(searchWord), make([][]string, 0, len(searchWord))
	for i := 1; i <= len(searchWord); i++ {
		searchWordBytes[i-1]++
		products = products[:sort.SearchStrings(products, string(searchWordBytes[:i]))]
		searchWordBytes[i-1]--
		products = products[sort.SearchStrings(products, searchWord[:i]):]
		if len(products) > 3 {
			result = append(result, products[:3])
		} else {
			result = append(result, products)
		}
	}
	return result
}
