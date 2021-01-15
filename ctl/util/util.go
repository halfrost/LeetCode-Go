package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

// LoadSolutionsDir define
func LoadSolutionsDir() ([]int, int) {
	files, err := ioutil.ReadDir("../leetcode/")
	if err != nil {
		fmt.Println(err)
	}
	solutionIds := []int{}
	for _, f := range files {
		if f.Name()[4] == '.' {
			tmp, err := strconv.Atoi(f.Name()[:4])
			if err != nil {
				fmt.Println(err)
			}
			solutionIds = append(solutionIds, tmp)
		}
	}
	sort.Ints(solutionIds)
	fmt.Printf("读取了 %v 道题的题解，当前目录下有 %v 个文件(可能包含 .DS_Store)，目录中有 %v 道题在尝试中\n", len(solutionIds), len(files), len(files)-len(solutionIds))
	return solutionIds, len(files) - len(solutionIds)
}

// WriteFile define
func WriteFile(fileName string, content []byte) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write file successful")
}

// BinarySearch define
func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
