package controller

import (
	"context"
	"path/filepath"

	"github.com/suzuki-shunsuke/clap/clap/install/extract"
)

type ParamsExtract struct {
	File   File
	Source string
}

func (ctrl Controller) Extract(ctx context.Context, params ParamsExtract) error {
	file := params.File
	dest := filepath.Dir(file.InstallPath)
	if fi, err := ctrl.FileChecker.Stat(file.InstallPath); err == nil {
		if fi.IsDir() {
			dest = file.InstallPath
		}
	}
	if err := ctrl.Extractor.Extract(extract.ParamsExtract{
		Source:      params.Source,
		Target:      file.Path,
		Destination: dest,
	}); err != nil {
		return err
	}
	return nil
}
