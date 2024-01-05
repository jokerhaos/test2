package main

import (
	"fmt"
	"image"
	"io"
	"net/http"
	"os"

	"github.com/corona10/goimagehash"
)

func main() {
	prefix := "85"
	suffix := "39"
	total := 100
	bufferCh := make(chan string, 10)
	file1Path := "./1.jpeg"
	file1, err := os.Open(file1Path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()
	img1, _, _ = image.Decode(file1)

	go func() {
		for i := 0; i < total; i++ {
			// 使用 strconv.FormatInt 将整数格式化为指定宽度的字符串，左侧补零
			middle := fmt.Sprintf("%05d", i)
			qq := prefix + middle + suffix
			bufferCh <- qq
		}
	}()

	// 循环缓冲写入数据
	for qq := range bufferCh {
		// 下载头像
		err := downPhoto(qq)
		if err == nil {
			// 头像对比
			b, _ := diff("./temp/" + qq + ".jpg")
			if b {
				fmt.Println("头像符合的QQ：" + qq)
				writeToFile("qq.txt", qq)
			}
		}
	}
	fmt.Println("程序执行完成...")
}

// 将符合条件的 QQ 号写入文件
func writeToFile(filename, qq string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(qq + "\n")
	if err != nil {
		return err
	}

	return nil
}

// 根据 QQ 号下载头像
func downPhoto(qq string) error {
	// 构建下载头像的 URL
	downPath := "http://q.qlogo.cn/headimg_dl?dst_uin=" + qq + "&spec=640&img_type=jpg"
	// 发起 HTTP GET 请求
	resp, err := http.Get(downPath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载头像失败，HTTP 状态码：%d", resp.StatusCode)
	}

	// 创建文件保存头像
	filePath := "./temp/" + qq + ".jpg"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应体内容复制到文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("头像下载成功：%s.jpg\n", qq)
	return nil
}

var img1 image.Image

// 图片对比
func diff(file2Path string) (bool, error) {
	file2, err := os.Open(file2Path)
	if err != nil {
		return false, err
	}
	defer file2.Close()

	img2, _, err := image.Decode(file2)
	if err != nil {
		return false, err
	}

	hash1, err := goimagehash.AverageHash(img1)
	if err != nil {
		return false, err
	}

	hash2, err := goimagehash.AverageHash(img2)
	if err != nil {
		return false, err
	}

	distance, err := hash1.Distance(hash2)
	if err != nil {
		return false, err
	}

	return distance < 100, nil
}
