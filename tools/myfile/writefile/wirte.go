package writefile

import (
	"strings"

	"github.com/pkg/errors"
)

type WriteInterface interface {
	Write(filePath string, data []byte) WriteResult
}

func Write(filePath string, data []byte) WriteResult {
	arr := strings.Split(filePath, ".")
	var fileSuffix string
	if len(arr) > 2 {
		fileSuffix = strings.Join(arr[len(arr)-1:], "")
	}
	switch fileSuffix {
	case "csv":
		return newCSVUtil().Write(filePath, data)
	case "json":
		return newJSONUtil().Write(filePath, data)
	}
	return newDefaultReadUtil().Write(filePath, data)
}

type defaultUtil struct {
	writeResult WriteResult
}

func newDefaultReadUtil() WriteInterface {
	return &defaultUtil{}
}

func (u *defaultUtil) Write(filePath string, data []byte) WriteResult {
	return WriteResult{
		Error: errors.New("un-support file suffix, " + filePath),
	}
}
