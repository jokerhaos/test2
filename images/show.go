package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// 指定要显示图片的文件夹路径
var imageFolderPath string = "./temp"

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/temp/", http.StripPrefix("/temp/", http.FileServer(http.Dir("temp"))))
	http.ListenAndServe(":9002", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// 获取文件夹下的所有文件
	files, err := os.ReadDir(imageFolderPath)
	if err != nil {
		http.Error(w, "Error reading image folder", http.StatusInternalServerError)
		return
	}

	// 模板定义
	const tpl = `
<!DOCTYPE html>
<html>
<head>
<style>
  .image-container {
    text-align: center;
    margin-bottom: 20px;
	width:200px
	;height:200px;
	float:left;
  }
</style>
</head>
<body>
  {{range .}}
    <div class="image-container">
      <img src="/temp/{{.ImageName}}" width="180" height="180">
      <b>{{.ImageName}}</b>
    </div>
  {{end}}
</body>
</html>
`

	// 解析模板
	tmpl, err := template.New("index").Parse(tpl)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// 构造图片信息的切片
	var images []ImageInfo
	for _, file := range files {
		if isImage(file.Name()) {
			images = append(images, ImageInfo{
				ImagePath: filepath.Join(imageFolderPath, file.Name()),
				ImageName: file.Name(),
			})
		}
	}

	// 执行模板
	w.WriteHeader(http.StatusOK) // 清除可能由于其他代码而设置的状态码
	err = tmpl.Execute(w, images)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

// ImageInfo 存储图片信息
type ImageInfo struct {
	ImagePath string
	ImageName string
}

// isImage 判断文件是否是图片
func isImage(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return true
	}
	return false
}
