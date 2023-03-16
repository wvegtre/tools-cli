package readfile

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type JSONUtil struct {
	readResult JSONReadResult
}

func NewJSONReadUtil() ReadInterface {
	return &JSONUtil{}
}

func (u *JSONUtil) Read(filePath string) ReadInterface {
	fs, err := os.Open(filePath)
	if err != nil {
		u.readResult = JSONReadResult{
			Error: err,
		}
		return u
	}
	defer fs.Close()

	byteValue, err := ioutil.ReadAll(fs)
	if err != nil {
		u.readResult = JSONReadResult{
			Error: err,
		}
		return u
	}
	u.readResult = JSONReadResult{
		Detail: byteValue,
	}
	return u
}

func (u *JSONUtil) Error() error {
	return u.readResult.Error
}

func (u *JSONUtil) GetFileContent(result interface{}) error {
	err := json.Unmarshal(u.readResult.Detail, result)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
