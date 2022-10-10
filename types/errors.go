package types

import "fmt"

type ErrFileNotFound struct {
	Path string
	err  string
}

type ErrDirNotFound struct {
	Path string
	err  string
}

func (e *ErrFileNotFound) Error() string {
	return fmt.Sprintf("file not found: %s", e.Path)
}

func (e *ErrDirNotFound) Error() string {
	return fmt.Sprintf("directory not found: %s", e.Path)
}
