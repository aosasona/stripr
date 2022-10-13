package types

import "fmt"

type FatalRuntimeError struct{}

type ErrFileNotFound struct {
	Path string
}

type ErrDirNotFound struct {
	Path string
}

type UnableToReadDir struct {
	Path string
}

type ErrFileNotText struct {
	Path string
}

func (e *FatalRuntimeError) Error() string {
	return "fatal runtime error, terminating..."
}

func (e *ErrFileNotFound) Error() string {
	return fmt.Sprintf("file not found: %s", e.Path)
}

func (e *ErrDirNotFound) Error() string {
	return fmt.Sprintf("directory not found: %s", e.Path)
}

func (e *UnableToReadDir) Error() string {
	return fmt.Sprintf("unable to read directory: %s", e.Path)
}

func (e *ErrFileNotText) Error() string {
	return fmt.Sprintf("file is not a text file: %s", e.Path)
}
