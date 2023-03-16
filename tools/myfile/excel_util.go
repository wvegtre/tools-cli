package myfile

//
//import (
//	"strconv"
//
//	"github.com/pkg/errors"
//	"github.com/tealeg/xlsx"
//)
//
//type ExcelUtil struct {
//	FilePath string
//	Write    []ExcelWriteParam
//}
//
//func (this ExcelUtil) StartWrite() error {
//	if len(this.Write) != 1 {
//		return errors.New("Write Param Invalid. Please Input Write Param Object.")
//	}
//	file := xlsx.NewFile()
//	sheetName := this.Write[0].SheetName
//	if sheetName == "" {
//		sheetName = DEFAULT_SHEET_1
//	}
//	if err := writeSheet(sheetName, file, this.Write[0]); err != nil {
//		return errors.WithStack(err)
//	}
//	if err := file.Save(this.FilePath); err != nil {
//		return errors.WithStack(err)
//	}
//	return nil
//}
//
//func (this ExcelUtil) StartMultipleWrite() error {
//	if len(this.Write) == 0 {
//		return errors.New("Write Param Invalid. Please Input Write Param Array.")
//	}
//	file := xlsx.NewFile()
//	var sheetName string
//	for i, v := range this.Write {
//		sheetName = v.SheetName
//		if sheetName == "" {
//			sheetName = "sheet" + strconv.Itoa(i+1)
//		}
//		if err := writeSheet(sheetName, file, v); err != nil {
//			return errors.WithStack(err)
//		}
//		if err := file.Save(this.FilePath); err != nil {
//			return errors.WithStack(err)
//		}
//	}
//	return nil
//}
//
//func (this ExcelUtil) StartRead() (map[string][]FileContentDetail, error) {
//	return readExcel(this.FilePath)
//}
//
//func (this ExcelUtil) StartReadHeads() ([]string, error) {
//	return nil, nil
//}
//
//func writeSheet(sheetName string, file *xlsx.File, wp ExcelWriteParam) error {
//	if sheetName == "" {
//		return errors.New("sheet name invalid. Please Input Sheet Name")
//	}
//	sheet, err := file.AddSheet(sheetName)
//	if err != nil {
//		return errors.WithStack(err)
//	}
//	// 添加头字段
//	if len(wp.Heads) > 0 {
//		row := sheet.AddRow()
//		for _, v := range wp.Heads {
//			cell := row.AddCell()
//			cell.Value = v
//		}
//	}
//	if len(wp.Datas) > 0 {
//		for _, d := range wp.Datas {
//			row := sheet.AddRow()
//			for _, v := range d {
//				cell := row.AddCell()
//				cell.Value = v
//			}
//		}
//	}
//	return nil
//}
//
//// 读取文件
//func readExcel(filePath string) (map[string][]FileContentDetail, error) {
//	// 打开文件
//	xlFile, err := xlsx.OpenFile(filePath)
//	if err != nil {
//		return nil, errors.WithStack(err)
//	}
//	if len(xlFile.Sheets) == 0 {
//		return nil, errors.New("Excel.Sheet is empty")
//	}
//	result := make(map[string][]FileContentDetail)
//	// 遍历sheet页读取
//	for _, sheet := range xlFile.Sheets {
//		if len(sheet.Rows) == 0 {
//			return nil, errors.New("sheet.Rows is empty")
//		}
//		// 记录单sheet页数据
//		// 拿到head信息
//		heads := getHeads(sheet.Rows[0])
//		// 详情数据不需要包含第一行
//		result[sheet.Name], err = getRowDetail(heads, sheet.Rows[1:])
//		if err != nil {
//			return nil, errors.WithStack(err)
//		}
//	}
//	return result, nil
//}
//
//func getHeads(row *xlsx.Row) []string {
//	heads := make([]string, 0, len(row.Cells))
//	for _, cell := range row.Cells {
//		if cell.String() == "" {
//			break
//		}
//		heads = append(heads, cell.String())
//	}
//	return heads
//}
//
//// 遍历每行的列读取，map存储
//// key: 字段名，value: 行上对应列的值
//func getRowDetail(heads []string, rows []*xlsx.Row) ([]FileContentDetail, error) {
//	result := make([]FileContentDetail, 0)
//
//	//遍历行读取
//	for _, row := range rows {
//		rowMap := make(map[string]string)
//		for i, cell := range row.Cells {
//			if i >= len(heads) {
//				return nil, errors.New("实际内容列不应该多于头字段列")
//			}
//			rowMap[genMapKey(heads[i])] = cell.String()
//		}
//		result = append(result, FileContentDetail{
//			Heads:   heads,
//			DataMap: rowMap,
//		})
//	}
//	return result, nil
//}
