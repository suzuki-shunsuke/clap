package fsys

import (
	"io"
	"os"
)

type FileCreator struct{}

type ParamsCreateFile struct {
	Reader io.Reader
	Source string
}

func (FileCreator) Create(params ParamsCreateFile) error {
	f, err := os.Create(params.Source)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, params.Reader); err != nil {
		return err
	}
	return nil
}
