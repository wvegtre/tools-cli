package writefile

import (
	"strings"

	"github.com/pkg/errors"
)

type WriteInterface interface {
	Write(filePath string, data []byte) WriteInterface
	Error() error
}

func Write(filePath string, data []byte) WriteInterface {
	arr := strings.Split(filePath, ".")
	var fileSuffix string
	if len(arr) > 2 {
		fileSuffix = strings.Join(arr[len(arr)-1:], "")
	}
	switch fileSuffix {
	case "csv":
		return NewCSVUtil().Write(filePath, data)
	case "json":
		return NewJSONUtil().Write(filePath, data)
	}
	return newDefaultReadUtil().Write(filePath, data)
}

type defaultUtil struct {
	writeResult WriteResult
}

func newDefaultReadUtil() WriteInterface {
	return &defaultUtil{}
}

func (u *defaultUtil) Write(filePath string, data []byte) WriteInterface {
	u.writeResult = WriteResult{
		Error: errors.New("un-support file suffix, " + filePath),
	}
	return u
}

func (u *defaultUtil) Error() error {
	return u.writeResult.Error
}
