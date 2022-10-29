package types

type Scanner struct {
	DirType DirType
	Path    string
	Config  map[string]interface{}
}

type (
	FilePath string
	FileName string
)

type ScanResult struct {
	Name        string
	Path        string
	Lines       [][]int
	HasComments bool
}

type DirType int

const (
	FILE DirType = iota
	DIRECTORY
)
