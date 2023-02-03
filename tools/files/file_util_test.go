package files

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestCheckFileType(t *testing.T) {
	filePath := "/Users/hb.li/Downloads/csv_import_test.csv"
	fileType := checkFileType(filePath)
	t.Log(fileType)
}

func TestCSVRead(t *testing.T) {
	filePath := "/Users/hb.li/Downloads/csv_import_test.csv"
	err := Read(filePath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("succeed")
}

type Demo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestZipRead(t *testing.T) {
	zipFile := "/Users/hb.li/Downloads/Archive.zip"
	zf, err := zip.OpenReader(zipFile)
	if err != nil {
		t.Error(err)
		return
	}
	defer zf.Close()

	for _, file := range zf.File {
		fmt.Printf("%s\n", file.Name)
		var dd Demo
		if err := readJsonFile(file, &dd); err != nil {
			t.Error(err)
			return
		}
		t.Log(fmt.Sprintf("name:%s, age:%d", dd.Name, dd.Age))
	}
	t.Log("succeed")
}

func readJsonFile(zf *zip.File, result interface{}) error {
	f, err := zf.Open()
	if err != nil {
		return errors.Unwrap(err)
	}
	defer f.Close()
	rb, err := io.ReadAll(f)
	if err != nil {
		return errors.Unwrap(err)
	}
	err = json.Unmarshal(rb, result)
	if err != nil {
		return errors.Unwrap(err)
	}
	return nil
}
