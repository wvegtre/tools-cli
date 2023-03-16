package myfile

import "strings"

const (
	DEFAULT_SHEET_1 = "sheet1"
)

type FileUtil interface {
	StartWrite() error
	StartMultipleWrite() error
	StartRead() (map[string][]FileContentDetail, error)
	StartReadHeads() ([]string, error)
}

func NewCsvUtil(filePath string, writeParam ...CsvWriteParam) FileUtil {
	return CsvUtil{
		FilePath: filePath,
		Write:    writeParam,
	}
}

func NewExcelUtil(filePath string, writeParam ...ExcelWriteParam) FileUtil {
	return ExcelUtil{
		FilePath: filePath,
		Write:    writeParam,
	}
}

func genMapKey(val string) string {
	vs := strings.Split(val, "_")
	var rest string
	for _, v := range vs {
		rest += strings.Title(v)
	}
	if rest == "" {
		rest = val
	}
	return rest
}
