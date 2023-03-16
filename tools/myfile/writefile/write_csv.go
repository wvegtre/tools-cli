package writefile

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type csvUtil struct {
	writeResult WriteResult
}

func newCSVUtil() WriteInterface {
	return &csvUtil{}
}

func (u *csvUtil) Write(filePath string, data []byte) WriteResult {
	parameter := &CSVWriteParameter{}
	err := json.Unmarshal(data, &parameter)
	if err != nil {
		return WriteResult{
			Error: errors.WithStack(err),
		}
	}
	err = genCSVFile(filePath, parameter)
	if err != nil {
		return WriteResult{
			Error: errors.WithStack(err),
		}
	}
	return WriteResult{}
}

func (u *csvUtil) Error() error {
	return u.writeResult.Error
}

func genCSVFile(filePath string, wd *CSVWriteParameter) error {
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
	if len(wd.Data) > 0 {
		for _, line := range wd.Data {
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
