package myfile

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

func (this CsvUtil) StartWrite() error {
	if len(this.Write) != 1 {
		return errors.New("Write Param Invalid. Please Input Write Param Object.")
	}
	return genCsvFile(this.FilePath, this.Write[0])
}

func (this CsvUtil) StartMultipleWrite() error {
	return errors.New("CSV 暂不支持多Sheet, 请改用其他模式。")
}

func (this CsvUtil) StartRead() (map[string][]FileContentDetail, error) {
	return readCsv(this.FilePath)
}

func (this CsvUtil) StartReadHeads() ([]string, error) {
	return readCsvHead(this.FilePath)
}

func genCsvFile(filePath string, wd CsvWriteParam) error {
	// 创建一个 tutorials.csv 文件
	csvFile, err := os.Create(filePath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer csvFile.Close()

	// 初始化一个 csv writer，并通过这个 writer 写入数据到 csv 文件
	writer := csv.NewWriter(csvFile)

	// 写入头数据
	if len(wd.Heads) > 0 {
		// 将切片类型行数据写入 csv 文件
		if err = writer.Write(wd.Heads); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(wd.Datas) > 0 {
		for _, line := range wd.Datas {
			// 将切片类型行数据写入 csv 文件
			if err = writer.Write(line); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	// 将 writer 缓冲中的数据都推送到 csv 文件，至此就完成了数据写入到 csv 文件
	writer.Flush()
	return nil
}

func readCsvHead(filePath string) ([]string, error) {
	// 打开文件
	fs, err := os.Open(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	var heads []string
	for i := 0; i < 1; i++ {
		row, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, errors.WithStack(err)
		}
		// 读取头字段信息
		if i == 0 {
			heads = row
			break
		}
	}
	return heads, err
}

func readCsv(filePath string) (map[string][]FileContentDetail, error) {

	// 打开文件
	fs, err := os.Open(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	var heads []string
	contents := make([]FileContentDetail, 0)
	for i := 0; ; i++ {
		row, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, errors.WithStack(err)
		}
		// 读取头字段信息
		if i == 0 {
			heads = row
			continue
		}
		// 具体内容按行读取
		rowMap := make(map[string]string)
		//遍历行读取
		for i, v := range row {
			if i >= len(heads) {
				return nil, errors.WithStack(err)
			}
			rowMap[genMapKey(heads[i])] = v
		}
		contents = append(contents, FileContentDetail{
			Heads:   heads,
			DataMap: rowMap,
		})
	}
	result := map[string][]FileContentDetail{
		DEFAULT_SHEET_1: contents,
	}
	return result, nil
}
