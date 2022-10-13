package types

type Scanner struct {
	DirType DirType
	Path    string
	Args    []string
}

type (
	FilePath string
	FileName string
)

type ScanResult struct {
	Name       FileName
	Path       FilePath
	Lines      []int
	LineCount  int
	HasComment bool
}

type DirType int

const (
	FILE DirType = iota
	DIRECTORY
)
