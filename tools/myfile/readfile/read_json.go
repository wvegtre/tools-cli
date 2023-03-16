package readfile

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type jsonUtil struct {
	readResult ReadResult
}

func newJSONReadUtil() ReadInterface {
	return &jsonUtil{}
}

func (u *jsonUtil) Read(filePath string) ReadResult {
	fs, err := os.Open(filePath)
	if err != nil {
		return ReadResult{
			Error: err,
		}
	}
	defer fs.Close()

	byteValue, err := ioutil.ReadAll(fs)
	if err != nil {
		return ReadResult{
			Error: err,
		}
	}
	return ReadResult{
		JSONDetail: byteValue,
	}
}

func (u *jsonUtil) Error() error {
	return u.readResult.Error
}

func (u *jsonUtil) GetFileContent(result interface{}) error {
	err := json.Unmarshal(u.readResult.JSONDetail, result)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
