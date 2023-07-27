package util

import (
	"github.com/samber/lo"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/threading"
)

type ExcelReader struct {
	file     *excelize.File
	DataChan chan [][]string
}

func NewExcelReader(file *excelize.File) *ExcelReader {
	return &ExcelReader{
		file:     file,
		DataChan: make(chan [][]string, 100),
	}
}

// start 从第几行开始读
// 每 n 份数量的数据发送到管道
func (r *ExcelReader) ReadAndSendToChannel(start int, num int) error {

	rows, err := r.file.Rows(r.file.GetSheetName(r.file.GetActiveSheetIndex()))
	if err != nil {
		return err
	}
	results, count, i := make([][]string, 0, num), 0, 0

	// 这里我用的gozero封装的协程，里面带了异常捕获
	threading.GoSafe(func() {
		defer close(r.DataChan)
		defer rows.Close()

		for rows.Next() {
			i++
			if start >= i {
				continue
			}
			row, err := rows.Columns()
			if err != nil {
				break
			}
			if len(row) > 0 {
				count++
				results = append(results, row)
			}
			if count%num == 0 {
				r.DataChan <- results
				count = 0
				results = make([][]string, 0, num)
			}
		}
		if count > 0 {
			r.DataChan <- results
		}
	})

	return nil
}

// start 从第几行开始读
// num 每 n 份数量的数据发送到管道
// columnsFilter 需要过滤的列
// predicate 过滤函数
func (r *ExcelReader) ReadAndSendToChannelFilter(start int, num int, columnsFilter []int, predicate func([]map[string]struct{}, []string) bool) error {

	rows, err := r.file.Rows(r.file.GetSheetName(r.file.GetActiveSheetIndex()))
	if err != nil {
		return err
	}
	results, count, i := make([][]string, 0, num), 0, 0

	dataAll := make([]map[string]struct{}, len(columnsFilter))
	// 这里我用的gozero封装的协程，里面带了异常捕获
	threading.GoSafe(func() {
		defer close(r.DataChan)
		defer rows.Close()

		for rows.Next() {
			i++
			if start >= i {
				continue
			}
			row, err := rows.Columns()
			if err != nil {
				break
			}

			// 长度判断 + 过滤函数
			if len(row) <= 0 || predicate(dataAll, row) {
				continue
			}

			// 加入map切片
			lo.ForEach[int](columnsFilter, func(item, _ int) {
				key := row[item]
				seen := dataAll[item]
				if seen == nil {
					seen = make(map[string]struct{})
				}
				seen[key] = struct{}{}
				dataAll[item] = seen
			})

			// 加入队列切片
			count++
			results = append(results, row)

			// 重置计算器和切片
			if count-num == 0 {
				r.DataChan <- results
				count = 0
				results = make([][]string, 0, num)
			}
		}
		if count > 0 {
			r.DataChan <- results
		}
	})

	return nil
}
