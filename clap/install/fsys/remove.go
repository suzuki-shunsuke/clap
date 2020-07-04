package fsys

import "os"

type FileRemover struct{}

func (FileRemover) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
