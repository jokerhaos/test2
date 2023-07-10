package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	concurrency := 10
	totalRows := 103
	totalOffset := 0
	level := 0

	// 打开文件并获取文件句柄
	fileHandle, err := os.OpenFile("output.csv", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	// 使用等待组来等待所有协程完成
	var wg sync.WaitGroup

	// 并发写入数据
	for totalOffset < totalRows {
		wg.Add(1)
		level++
		level := level
		// 计算当前协程需要写入的起始行和结束行
		startRow := totalOffset
		endRow := startRow + concurrency

		// 最后一个协程处理余下的行数
		if endRow > totalRows {
			endRow = totalRows
		}
		// 启动协程并发写入
		go func() {
			defer wg.Done()

			// 获取数据
			data := fetchData(startRow, endRow)

			// 单行预估数值
			rowSize := len(data[0])
			// X行数据大小 = 单行预估数值 * X + 扩容预估值
			estimateSize := (rowSize * concurrency) + 100
			// 当前数据位置 = X行数据大小 * level - 最后一行的多余的预估值
			thisPtr := estimateSize*level - ((concurrency - len(data)) * rowSize)
			// 移动指针
			if startRow == 0 {
				thisPtr = 0
			}
			offset, err := fileHandle.Seek(int64(thisPtr), 0)

			fmt.Println(startRow, endRow, level, thisPtr, data)

			if err != nil {
				panic(err)
			}

			// 写入数据到文件指针位置
			for _, row := range data {
				context := []byte(fmt.Sprintf("%s\n", row))
				_, err := fileHandle.WriteAt(context, offset)
				if err != nil {
					panic(err)
				}
				offset += int64(len(context))

				// fmt.Println(offset, string(context))

			}
		}()

		totalOffset += concurrency

	}

	wg.Wait()

	fmt.Println("success")
}

func fetchData(startRow, endRow int) []string {
	// 这里可以替换为自己的数据获取逻辑，根据startRow和endRow获取对应的数据
	// 例如，从数据库查询指定范围的数据
	data := make([]string, endRow-startRow)
	for i := startRow; i < endRow; i++ {
		data[i-startRow] = fmt.Sprintf("data%d", i+1)
	}
	return data
}
