package myfile

//
//func TestReadCSV(t *testing.T) {
//
//	filePath := "/Users/hb.li/Documents/AfterShip/db/cn-p-core/order/orders_cancel_20210401.csv"
//	rest, er := NewCsvUtil(filePath).StartRead()
//	if er != nil {
//		t.Error("Write CSV Error, ", er)
//	}
//	t.Log("csv read success")
//	for k, v := range rest {
//		t.Log(k, ", 总数: ", len(v))
//		for _, content := range v {
//			t.Log("DataMap --> ", content.DataMap)
//		}
//	}
//}
//
//func TestWriteCSV(t *testing.T) {
//	saveFilePath := "/Users/hb.li/Documents/AfterShip/output"
//	saveFileName := "/test_save.csv"
//	wp := CsvWriteParam{
//		Heads: []string{"r1", "r2", "r3", "r4", "r5", "r6", "r7"},
//		Data: [][]string{
//			{"v11", "v12", "v13", "v14", "v15", "v16", "v17"},
//			{"v21", "v22", "v23", "v24", "v25", "v26", "v27"},
//			{"v31", "v32", "v33", "v34", "v35", "v36", "v37"},
//			{"v41", "v42", "v43", "v44", "v45", "v46", "v47"},
//		},
//	}
//	if er := NewCsvUtil(saveFilePath+saveFileName, wp).StartWrite(); er != nil {
//		t.Error("write csv error, ", er)
//		return
//	}
//	t.Log("csv write success, filePath: ", saveFilePath+saveFileName)
//}
//
//func TestReadExcel(t *testing.T) {
//	filePath := "/Users/hb.li/Downloads/Dropshipping_Worker_Pub_Sub.xlsx"
//	rest, er := NewExcelUtil(filePath).StartRead()
//	if er != nil {
//		t.Error("Write Excel Error, ", er)
//	}
//	t.Log("excel read success")
//	for k, v := range rest {
//		t.Log(k, ", 总数: ", len(v))
//		for _, content := range v {
//			t.Log("DataMap --> ", content.DataMap)
//		}
//	}
//}
//
//func TestWriteExcel(t *testing.T) {
//	saveFilePath := "/Users/hb.li/Documents/AfterShip/output"
//	saveFileName := "/test_save.xlsx"
//	wp := ExcelWriteParam{
//		Heads: []string{"r1", "r2", "r3", "r4", "r5", "r6", "r7"},
//		Data: [][]string{
//			{"v11", "v12", "v13", "v14", "v15", "v16", "v17"},
//			{"v21", "v22", "v23", "v24", "v25", "v26", "v27"},
//			{"v31", "v32", "v33", "v34", "v35", "v36", "v37"},
//			{"v41", "v42", "v43", "v44", "v45", "v46", "v47"},
//		},
//	}
//	if er := NewExcelUtil(saveFilePath+saveFileName, wp).StartWrite(); er != nil {
//		t.Error("write csv error, ", er)
//		return
//	}
//	t.Log("csv write success, filePath: ", saveFilePath+saveFileName)
//}
//
//func TestMultipleWriteExcel(t *testing.T) {
//	saveFilePath := "/Users/hb.li/Documents/AfterShip/output"
//	saveFileName := "/test_save_multiple.xlsx"
//	wps := make([]ExcelWriteParam, 0)
//	wps = append(wps, ExcelWriteParam{
//		SheetName: "sheet-1",
//		Heads:     []string{"r1", "r2", "r3", "r4", "r5", "r6", "r7"},
//		Data: [][]string{
//			{"v11", "v12", "v13", "v14", "v15", "v16", "v17"},
//			{"v21", "v22", "v23", "v24", "v25", "v26", "v27"},
//			{"v31", "v32", "v33", "v34", "v35", "v36", "v37"},
//			{"v41", "v42", "v43", "v44", "v45", "v46", "v47"},
//		},
//	})
//	wps = append(wps, ExcelWriteParam{
//		SheetName: "sheet-2",
//		Heads:     []string{"rr1", "rr2", "rr3", "rr4", "rr5", "rr6", "rr7"},
//		Data: [][]string{
//			{"vv11", "v12", "v13", "v14", "v15", "v16", "v17"},
//			{"vv21", "v22", "v23", "v24", "v25", "v26", "v27"},
//			{"vv31", "v32", "v33", "v34", "v35", "v36", "v37"},
//			{"vv41", "v42", "v43", "v44", "v45", "v46", "v47"},
//		},
//	})
//	if er := NewExcelUtil(saveFilePath+saveFileName, wps...).StartMultipleWrite(); er != nil {
//		t.Error("write csv error, ", er)
//		return
//	}
//	t.Log("csv write success, filePath: ", saveFilePath+saveFileName)
//}
