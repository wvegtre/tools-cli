package readfile

type ReadResult struct {
	Error      error
	CSVDetail  *FileContentDetail
	JSONDetail []byte
}

type FileContentDetail struct {
	SheetName string
	Heads     []string
	Data      [][]string
}
