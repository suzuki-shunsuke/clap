package controller

import (
	"context"
	"path/filepath"

	"github.com/suzuki-shunsuke/clap/pkg/extract"
)

type ParamsExtract struct {
	File   File
	Source string
}

func (ctrl Controller) Extract(ctx context.Context, baseDir string, params ParamsExtract) error {
	file := params.File
	dest := file.InstallPath
	if err := ctrl.Mkdir.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	if fi, err := ctrl.FileChecker.Stat(file.InstallPath); err == nil {
		if fi.IsDir() {
			dest = filepath.Join(file.InstallPath, filepath.Base(file.Path))
		}
	}
	if err := ctrl.Extractor.Extract(extract.ParamsExtract{
		Source:      params.Source,
		Target:      file.Path,
		Destination: baseDir,
	}); err != nil {
		return err
	}
	if err := ctrl.FileRemover.RemoveAll(dest); err != nil {
		return err
	}
	if err := ctrl.FileRenamer.Rename(filepath.Join(baseDir, file.Path), dest); err != nil {
		return err
	}
	return nil
}
