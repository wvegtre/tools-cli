package writefile

type WriteResult struct {
	Error error
}

type CSVWriteParameter struct {
	Heads []string
	Data  [][]string
}
