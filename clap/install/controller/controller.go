package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
)

type File struct {
	Path        string
	InstallPath string
}

type ParamsRun struct {
	URL   *url.URL
	Files []File
}

var ErrEmptyFileName error = errors.New("file name is empty")

func (ctrl Controller) Run(ctx context.Context, params ParamsRun) error {
	tempDir, err := ctrl.TempDir.Create()
	if err != nil {
		return fmt.Errorf("failed to create a temporal directory: %w", err)
	}

	defer func() {
		if err := ctrl.FileRemover.RemoveAll(tempDir); err != nil {
			log.Printf("failed to remove a temporary directory %s: %v", tempDir, err)
		}
	}()

	fileName := filepath.Base(params.URL.Path)
	if fileName == "" {
		return ErrEmptyFileName
	}
	src := filepath.Join(tempDir, fileName)

	if err := ctrl.Download(ctx, ParamsDownload{
		URL:    params.URL,
		Source: src,
		Dir:    tempDir,
	}); err != nil {
		return err
	}

	var result error
	for _, file := range params.Files {
		if err := ctrl.Extract(ctx, ParamsExtract{
			File:   file,
			Source: src,
		}); err != nil {
			result = multierror.Append(result, err)
			continue
		}
	}
	return result
}
