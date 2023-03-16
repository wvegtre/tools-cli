package writefile

import (
	"encoding/json"
	"os"
)

type jsonUtil struct {
	writeResult WriteResult
}

func newJSONUtil() WriteInterface {
	return &jsonUtil{}
}

func (u *jsonUtil) Write(filePath string, data []byte) WriteResult {
	fs, err := os.Create(filePath)
	if err != nil {
		return WriteResult{
			Error: err,
		}
	}
	defer fs.Close()
	encoder := json.NewEncoder(fs)
	err = encoder.Encode(data)
	if err != nil {
		return WriteResult{
			Error: err,
		}
	}
	return WriteResult{}
}

func (u *jsonUtil) Error() error {
	return u.writeResult.Error
}
