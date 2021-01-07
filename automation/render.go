package main

import (
	"encoding/json"
	"fmt"
	m "github.com/halfrost/LeetCode-Go/automation/models"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func main() {
	resp, err := http.Get("https://leetcode.com/api/problems/all/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var result []m.StatStatusPairs
	var lpa m.LeetCodeProblemAll
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &lpa)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = lpa.StatStatusPairs
	//fmt.Println(result)
	mdrows := []m.Mdrow{}
	for i := 0; i < len(result); i++ {
		mdrows = append(mdrows, convertModel(result[i]))
	}
	sort.Sort(m.SortByQuestionId(mdrows))
	res, _ := json.Marshal(mdrows)
	write(res)
	//fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
}

func write(content []byte) {
	file, err := os.OpenFile("leetcode_problem", os.O_RDWR|os.O_CREATE, 0777)
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

func convertModel(ssp m.StatStatusPairs) m.Mdrow {
	res := m.Mdrow{}
	res.FrontendQuestionId = strconv.FormatInt(int64(ssp.Stat.FrontendQuestionId), 10)
	res.QuestionTitle = ssp.Stat.QuestionTitle
	res.QuestionTitleSlug = ssp.Stat.QuestionTitleSlug
	// res.SolutionPath
	res.Acceptance = fmt.Sprintf("%.1f%%", (ssp.Stat.TotalAcs/ssp.Stat.TotalSubmitted)*100)
	res.Difficulty = m.DifficultyMap[ssp.Difficulty.Level]
	res.Frequency = fmt.Sprintf("%f", ssp.Frequency)
	return res
}
