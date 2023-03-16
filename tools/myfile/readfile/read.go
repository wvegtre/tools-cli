package readfile

import (
	"strings"

	"github.com/pkg/errors"
)

type ReadInterface interface {
	Read(filePath string) ReadResult
}

func Read(filePath string) ReadResult {
	arr := strings.Split(filePath, ".")
	var fileSuffix string
	if len(arr) > 2 {
		fileSuffix = strings.Join(arr[len(arr)-1:], "")
	}
	switch fileSuffix {
	case "csv":
		return newCSVReadUtil().Read(filePath)
	case "json":
		return newJSONReadUtil().Read(filePath)
	}
	return newDefaultReadUtil().Read(filePath)
}

type defaultReadUtil struct {
}

func newDefaultReadUtil() ReadInterface {
	return &defaultReadUtil{}
}

func (u *defaultReadUtil) Read(filePath string) ReadResult {
	return ReadResult{
		Error: errors.New("un-support file suffix, " + filePath),
	}
}
