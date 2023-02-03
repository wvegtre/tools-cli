package files

import (
	"errors"
	"os"
	"strings"
)

const (
	FileTypeUnKnown = iota
	FileTypeCSV
	FileTypeZIP
)

type fileUtilInterface interface {
	Read(f *os.File) error
	Write() error
}

func Read(file string) error {
	fileType := checkFileType(file)
	h, ok := _fileHandles[fileType]
	if !ok {
		return errors.New("un-support file suffix")
	}
	// open file
	f, err := os.Open(file)
	if err != nil {
		return errors.Unwrap(err)
	}
	defer f.Close()

	err = h.Read(f)
	if err != nil {
		return errors.Unwrap(err)
	}
	return nil
}

func Write() error {
	return nil
}

func checkFileType(file string) int {
	arr := strings.Split(file, ".")
	length := len(arr)
	var fileSuffix string
	if length > 2 {
		fileSuffix = arr[length-1 : length][0]
	}
	switch fileSuffix {
	case "csv":
		return FileTypeCSV
	}
	return FileTypeUnKnown
}

var _fileHandles map[int]fileUtilInterface

func register(fileType int, impl fileUtilInterface) {
	if _fileHandles == nil {
		_fileHandles = make(map[int]fileUtilInterface)
	}
	_fileHandles[fileType] = impl
}
