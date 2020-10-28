package fsys

import (
	"fmt"
	"io"
	"os"
)

type FileRenamer struct{}

func (FileRenamer) Rename(src, dest string) error {
	if err := os.Rename(src, dest); err != nil {
		return create(src, dest)
	}
	return nil
}

func create(src, dest string) error {
	inputFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open a source file "+src+": %w", err)
	}
	closed := false
	defer func() {
		if !closed {
			inputFile.Close()
		}
	}()
	outputFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to open a dest file "+dest+": %w", err)
	}
	defer outputFile.Close()
	if _, err := io.Copy(outputFile, inputFile); err != nil {
		return fmt.Errorf("failed to copy src to dest: %w", err)
	}
	inputFile.Close()
	closed = true
	if err := os.Remove(src); err != nil {
		return fmt.Errorf("failed to remove src "+src+": %w", err)
	}
	return nil
}
