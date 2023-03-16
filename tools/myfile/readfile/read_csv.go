package readfile

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

type CSVUtil struct {
	readResult CSVReadResult
}

func NewCSVReadUtil() ReadInterface {
	return &CSVUtil{}
}

func (u *CSVUtil) Read(filePath string) ReadInterface {
	content, err := readCsv(filePath)
	if err != nil {
		u.readResult = CSVReadResult{
			Error: errors.WithStack(err),
		}
		return u
	}
	u.readResult = CSVReadResult{
		Detail: content,
	}
	return u
}

func (u *CSVUtil) Error() error {
	return u.readResult.Error
}

func (u *CSVUtil) GetFileContent() *FileContentDetail {
	return u.readResult.Detail
}

func readCsv(filePath string) (*FileContentDetail, error) {
	// 打开文件
	fs, err := os.Open(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	var headers []string
	var data [][]string
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
			headers = row
			continue
		}
		// 保存每一行的数据，下标 + 1 则可以从 headers 里边匹配到是哪一列的数据
		data = append(data, row)
	}
	return &FileContentDetail{
		Heads: headers,
		Data:  data,
	}, nil
}
