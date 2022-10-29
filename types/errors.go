package types

import "fmt"

type (
	FatalRuntimeError struct{}
	ErrNoCommand      struct{}
	ErrInvalidCommand struct{ Command string }
	ErrFileNotFound   struct{ Path string }
	ErrDirNotFound    struct{ Path string }
	UnableToReadDir   struct{ Path string }
	ErrFileNotText    struct{ Path string }
	ErrFileIgnored    struct{ Path string }
)

func (e *FatalRuntimeError) Error() string {
	return "fatal runtime error, exiting..."
}

func (e *ErrNoCommand) Error() string {
	return "no command provided, exiting..."
}

func (e *ErrInvalidCommand) Error() string {
	return fmt.Sprintf("invalid command: %s", e.Command)
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

func (e *ErrFileIgnored) Error() string {
	return fmt.Sprintf("file ignored: %s", e.Path)
}
