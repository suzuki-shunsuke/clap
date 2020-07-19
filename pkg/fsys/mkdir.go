package fsys

import (
	"os"
)

type Mkdir struct{}

func (Mkdir) MkdirAll(path string, mode os.FileMode) error {
	return os.MkdirAll(path, mode)
}
