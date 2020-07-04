package fsys

import "os"

type FileChecker struct{}

func (FileChecker) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}
