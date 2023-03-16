package writefile

import (
	"encoding/json"
	"os"
)

type JSONUtil struct {
	writeResult WriteResult
}

func NewJSONUtil() WriteInterface {
	return &JSONUtil{}
}

func (u *JSONUtil) Write(filePath string, data []byte) WriteInterface {
	fs, err := os.Create(filePath)
	if err != nil {
		u.writeResult = WriteResult{
			Error: err,
		}
		return u
	}
	defer fs.Close()
	encoder := json.NewEncoder(fs)
	err = encoder.Encode(data)
	if err != nil {
		u.writeResult = WriteResult{
			Error: err,
		}
		return u
	}
	return u
}

func (u *JSONUtil) Error() error {
	return u.writeResult.Error
}
