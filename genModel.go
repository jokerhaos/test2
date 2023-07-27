package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// 定义模板参数结构体
type TemplateData struct {
	Package string
	Model   string
}

func main() {
	// 读取模板文件内容
	templateFile := "./template/model.tmpl"
	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	// 定义模板参数
	data := TemplateData{
		Package: "model",
		Model:   "User",
	}

	// 创建模板并解析
	tmplParsed := template.Must(template.New("customTemplate").Parse(string(templateContent)))

	// 使用模板和参数生成代码
	var outputCode strings.Builder
	err = tmplParsed.Execute(&outputCode, data)
	if err != nil {
		panic(err)
	}

	// 设置生成代码的路径和文件名
	outputPath := "./model/" // 替换成您希望的输出路径
	outputFileName := "xxx.gen.go"

	// 将生成的代码写入文件
	outputFile := outputPath + outputFileName
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(outputCode.String())
	fmt.Printf("Code generated and saved to '%s'\n", outputFile)
}
