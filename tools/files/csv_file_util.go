package files

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

type csvFileUtil string

func init() {
	var cfu csvFileUtil
	register(FileTypeCSV, cfu)
}

func (u csvFileUtil) Read(f *os.File) error {
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Unwrap(err)
		}
		for i, v := range rec {
			log.Println(i, "-->", v)
		}
	}
	return nil
}

func (u csvFileUtil) Write() error {
	return nil
}
