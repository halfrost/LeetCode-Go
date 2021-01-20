package main

import (
	"bytes"
	"fmt"
	m "github.com/halfrost/LeetCode-Go/ctl/models"
	"github.com/halfrost/LeetCode-Go/ctl/util"
	"html/template"
	"io/ioutil"
	"os"
)

func makeReadmeFile(mdrows m.Mdrows) {
	file := "./README.md"
	os.Remove(file)
	var b bytes.Buffer
	tmpl := template.Must(template.New("readme").Parse(readTMPL("template.markdown")))
	err := tmpl.Execute(&b, mdrows)
	if err != nil {
		fmt.Println(err)
	}
	// 保存 README.md 文件
	util.WriteFile(file, b.Bytes())
}

func readTMPL(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}
