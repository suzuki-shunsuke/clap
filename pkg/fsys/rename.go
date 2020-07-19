package fsys

import "os"

type FileRenamer struct{}

func (FileRenamer) Rename(src, dest string) error {
	return os.Rename(src, dest)
}
