package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/halfrost/LeetCode-Go/ctl/util"
)

// Graphql define
type Graphql struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		TitleSlug string `json:"titleSlug"`
	} `json:"variables"`
	Query string `json:"query"`
}

// GraphQLResp define
type GraphQLResp struct {
	Data struct {
		TopicTag       TopicTag       `json:"topicTag"`
		FavoritesLists FavoritesLists `json:"favoritesLists"`
	} `json:"data"`
}

// TopicTag define
type TopicTag struct {
	Name           string     `json:"name"`
	TranslatedName string     `json:"translatedName"`
	Slug           string     `json:"slug"`
	Questions      []Question `json:"questions"`
	Frequencies    string     `json:"frequencies"`
	Typename       string     `json:"__typename"`
}

// Question define
type Question struct {
	Status             string      `json:"status"`
	QuestionID         string      `json:"questionId"`
	QuestionFrontendID string      `json:"questionFrontendId"`
	Title              string      `json:"title"`
	TitleSlug          string      `json:"titleSlug"`
	TranslatedTitle    string      `json:"translatedTitle"`
	Stats              string      `json:"stats"`
	Difficulty         string      `json:"difficulty"`
	TopicTags          []TopicTags `json:"topicTags"`
	CompanyTags        interface{} `json:"companyTags"`
	Typename           string      `json:"__typename"`
}

// TopicTags define
type TopicTags struct {
	Name           string `json:"name"`
	TranslatedName string `json:"translatedName"`
	Slug           string `json:"slug"`
	Typename       string `json:"__typename"`
}

func (q Question) generateTagStatus() (TagStatus, error) {
	var ts TagStatus
	err := json.Unmarshal([]byte(q.Stats), &ts)
	if err != nil {
		fmt.Println(err)
		return ts, err
	}
	return ts, nil
}

// TagStatus define
type TagStatus struct {
	TotalAccepted      string `json:"totalAccepted"`
	TotalSubmission    string `json:"totalSubmission"`
	TotalAcceptedRaw   int32  `json:"totalAcceptedRaw"`
	TotalSubmissionRaw int32  `json:"totalSubmissionRaw"`
	AcRate             string `json:"acRate"`
}

// ConvertMdModelFromQuestions define
func ConvertMdModelFromQuestions(questions []Question) []Mdrow {
	mdrows := []Mdrow{}
	for _, question := range questions {
		res := Mdrow{}
		v, _ := strconv.Atoi(question.QuestionFrontendID)
		res.FrontendQuestionID = int32(v)
		res.QuestionTitle = strings.TrimSpace(question.Title)
		res.QuestionTitleSlug = strings.TrimSpace(question.TitleSlug)
		q, err := question.generateTagStatus()
		if err != nil {
			fmt.Println(err)
		}
		res.Acceptance = q.AcRate
		res.Difficulty = question.Difficulty
		mdrows = append(mdrows, res)
	}
	return mdrows
}

// TagList define
type TagList struct {
	FrontendQuestionID int32  `json:"question_id"`
	QuestionTitle      string `json:"question__title"`
	SolutionPath       string `json:"solution_path"`
	Acceptance         string `json:"acceptance"`
	Difficulty         string `json:"difficulty"`
	TimeComplexity     string `json:"time_complexity"`
	SpaceComplexity    string `json:"space_complexity"`
	Favorite           string `json:"favorite"`
}

// | 0001 | Two Sum  | [Go]({{< relref "/ChapterFour/0001.Two-Sum.md" >}})| Easy | O(n)| O(n)|❤️|50%|
func (t TagList) tableLine() string {
	return fmt.Sprintf("|%04d|%v|%v|%v|%v|%v|%v|%v|\n", t.FrontendQuestionID, t.QuestionTitle, t.SolutionPath, t.Difficulty, t.TimeComplexity, t.SpaceComplexity, t.Favorite, t.Acceptance)
}

func standardizedTitle(orig string, frontendQuestionID int32) string {
	s0 := strings.TrimSpace(orig)
	s1 := strings.Replace(s0, " ", "-", -1)
	s2 := strings.Replace(s1, "'", "", -1)
	s3 := strings.Replace(s2, "%", "", -1)
	s4 := strings.Replace(s3, "(", "", -1)
	s5 := strings.Replace(s4, ")", "", -1)
	s6 := strings.Replace(s5, ",", "", -1)
	s7 := strings.Replace(s6, "?", "", -1)
	count := 0
	// 去掉 --- 这种情况，这种情况是由于题目标题中包含 - ，左右有空格，左右一填充，造成了 ---，3 个 -
	for i := 0; i < len(s7)-2; i++ {
		if s7[i] == '-' && s7[i+1] == '-' && s7[i+2] == '-' {
			fmt.Printf("【需要修正 --- 的标题是 %v】\n", fmt.Sprintf("%04d.%v", int(frontendQuestionID), s7))
			s7 = s7[:i+1] + s7[i+3:]
			count++
		}
	}
	if count > 0 {
		fmt.Printf("总共修正了 %v 个标题\n", count)
	}
	// 去掉 -- 这种情况，这种情况是由于题目标题中包含负号 -
	for i := 0; i < len(s7)-2; i++ {
		if s7[i] == '-' && s7[i+1] == '-' {
			fmt.Printf("【需要修正 -- 的标题是 %v】\n", fmt.Sprintf("%04d.%v", int(frontendQuestionID), s7))
			s7 = s7[:i+1] + s7[i+2:]
			count++
		}
	}
	if count > 0 {
		fmt.Printf("总共修正了 %v 个标题\n", count)
	}
	return s7
}

// GenerateTagMdRows define
func GenerateTagMdRows(solutionIds []int, metaMap map[int]TagList, mdrows []Mdrow, internal bool) []TagList {
	tl := []TagList{}
	for _, row := range mdrows {
		if util.BinarySearch(solutionIds, int(row.FrontendQuestionID)) != -1 {
			tmp := TagList{}
			tmp.FrontendQuestionID = row.FrontendQuestionID
			tmp.QuestionTitle = strings.TrimSpace(row.QuestionTitle)
			s7 := standardizedTitle(row.QuestionTitle, row.FrontendQuestionID)
			if internal {
				tmp.SolutionPath = fmt.Sprintf("[Go]({{< relref \"/ChapterFour/%v/%v.md\" >}})", util.GetChpaterFourFileNum(int(row.FrontendQuestionID)), fmt.Sprintf("%04d.%v", int(row.FrontendQuestionID), s7))
			} else {
				tmp.SolutionPath = fmt.Sprintf("[Go](https://books.halfrost.com/leetcode/ChapterFour/%v/%v)", util.GetChpaterFourFileNum(int(row.FrontendQuestionID)), fmt.Sprintf("%04d.%v", int(row.FrontendQuestionID), s7))
			}
			tmp.Acceptance = row.Acceptance
			tmp.Difficulty = row.Difficulty
			tmp.TimeComplexity = metaMap[int(row.FrontendQuestionID)].TimeComplexity
			tmp.SpaceComplexity = metaMap[int(row.FrontendQuestionID)].SpaceComplexity
			tmp.Favorite = metaMap[int(row.FrontendQuestionID)].Favorite
			tl = append(tl, tmp)
		}
	}
	return tl
}

// TagLists define
type TagLists struct {
	TagLists []TagList
}

// | No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
// |:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
func (tls TagLists) table() string {
	res := "| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |\n"
	res += "|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |\n"
	for _, p := range tls.TagLists {
		res += p.tableLine()
	}
	// 加这一行是为了撑开整个表格
	res += "|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|"
	return res
}

// AvailableTagTable define
func (tls TagLists) AvailableTagTable() string {
	return tls.table()
}

// FavoritesLists define
type FavoritesLists struct {
	PublicFavorites  []int `json:"publicFavorites"`
	PrivateFavorites []struct {
		IDHash           string `json:"idHash"`
		ID               string `json:"id"`
		Name             string `json:"name"`
		IsPublicFavorite bool   `json:"isPublicFavorite"`
		ViewCount        int    `json:"viewCount"`
		Creator          string `json:"creator"`
		IsWatched        bool   `json:"isWatched"`
		Questions        []struct {
			QuestionID string `json:"questionId"`
			Title      string `json:"title"`
			TitleSlug  string `json:"titleSlug"`
			Typename   string `json:"__typename"`
		} `json:"questions"`
		Typename string `json:"__typename"`
	} `json:"privateFavorites"`
	Typename string `json:"__typename"`
}

// Gproblem define
type Gproblem struct {
	QuestionID            string        `json:"questionId"`
	QuestionFrontendID    string        `json:"questionFrontendId"`
	BoundTopicID          int           `json:"boundTopicId"`
	Title                 string        `json:"title"`
	TitleSlug             string        `json:"titleSlug"`
	Content               string        `json:"content"`
	TranslatedTitle       string        `json:"translatedTitle"`
	TranslatedContent     string        `json:"translatedContent"`
	IsPaidOnly            bool          `json:"isPaidOnly"`
	Difficulty            string        `json:"difficulty"`
	Likes                 int           `json:"likes"`
	Dislikes              int           `json:"dislikes"`
	IsLiked               interface{}   `json:"isLiked"`
	SimilarQuestions      string        `json:"similarQuestions"`
	Contributors          []interface{} `json:"contributors"`
	LangToValidPlayground string        `json:"langToValidPlayground"`
	TopicTags             []struct {
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		TranslatedName string `json:"translatedName"`
		Typename       string `json:"__typename"`
	} `json:"topicTags"`
	CompanyTagStats interface{}    `json:"companyTagStats"`
	CodeSnippets    []GcodeSnippet `json:"codeSnippets"`
	Stats           string         `json:"stats"`
	Hints           []interface{}  `json:"hints"`
	Solution        interface{}    `json:"solution"`
	Status          interface{}    `json:"status"`
	SampleTestCase  string         `json:"sampleTestCase"`
	MetaData        string         `json:"metaData"`
	JudgerAvailable bool           `json:"judgerAvailable"`
	JudgeType       string         `json:"judgeType"`
	MysqlSchemas    []interface{}  `json:"mysqlSchemas"`
	EnableRunCode   bool           `json:"enableRunCode"`
	EnableTestMode  bool           `json:"enableTestMode"`
	EnvInfo         string         `json:"envInfo"`
	Typename        string         `json:"__typename"`
}

// Gstat define
type Gstat struct {
	TotalAcs            int    `json:"total_acs"`
	QuestionTitle       string `json:"question__title"`
	IsNewQuestion       bool   `json:"is_new_question"`
	QuestionArticleSlug string `json:"question__article__slug"`
	TotalSubmitted      int    `json:"total_submitted"`
	FrontendQuestionID  int    `json:"frontend_question_id"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionArticleLive bool   `json:"question__article__live"`
	QuestionHide        bool   `json:"question__hide"`
	QuestionID          int    `json:"question_id"`
}

// GcodeSnippet define
type GcodeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
	Typename string `json:"__typename"`
}
