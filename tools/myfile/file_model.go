package myfile

// 读取结果
type FileContentDetail struct {
	//SheetName string
	Heads   []string
	DataMap map[string]string
}

type ExcelUtil struct {
	FilePath string
	Write    []ExcelWriteParam
}

type ExcelWriteParam struct {
	SheetName string
	Heads     []string
	Datas     [][]string
}

type ExcelRowWriteDetail struct {
	Cells []ExcelCellWriteDetail
}

// key为列位置，value为实际值，方便写的时候定位
type ExcelCellWriteDetail struct {
	Cell map[int]string
}

type CsvUtil struct {
	FilePath string
	Write    []CsvWriteParam
}

type CsvWriteParam struct {
	Heads []string
	Datas [][]string
}
