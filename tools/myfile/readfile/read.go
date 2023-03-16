package readfile

import (
	"strings"

	"github.com/pkg/errors"
)

type ReadInterface interface {
	Read(filePath string) ReadInterface
	Error() error
}

func Read(filePath string) ReadInterface {
	arr := strings.Split(filePath, ".")
	var fileSuffix string
	if len(arr) > 2 {
		fileSuffix = strings.Join(arr[len(arr)-1:], "")
	}
	switch fileSuffix {
	case "csv":
		return NewCSVReadUtil().Read(filePath)
	case "json":
		return NewJSONReadUtil().Read(filePath)
	}
	return newDefaultReadUtil().Read(filePath)
}

type defaultReadUtil struct {
	readResult defaultReadResult
}

func newDefaultReadUtil() ReadInterface {
	return &defaultReadUtil{}
}

func (u *defaultReadUtil) Read(filePath string) ReadInterface {
	u.readResult = defaultReadResult{
		Error: errors.New("un-support file suffix, " + filePath),
	}
	return u
}

func (u *defaultReadUtil) Error() error {
	return u.readResult.Error
}
