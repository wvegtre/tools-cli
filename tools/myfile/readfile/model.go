package readfile

type defaultReadResult struct {
	Error error
}

type CSVReadResult struct {
	Error  error
	Detail *FileContentDetail
}

type FileContentDetail struct {
	SheetName string
	Heads     []string
	Data      [][]string
}

type JSONReadResult struct {
	Error  error
	Detail []byte
}
