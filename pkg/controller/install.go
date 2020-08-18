package controller

import (
	"context"
	"path/filepath"
)

type ParamsInstall struct {
	File   File
	Source string
}

func (ctrl Controller) Install(ctx context.Context, baseDir string, params ParamsInstall) error {
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
	if err := ctrl.FileRemover.RemoveAll(dest); err != nil {
		return err
	}
	if err := ctrl.FileRenamer.Rename(filepath.Join(baseDir, file.Path), dest); err != nil {
		return err
	}
	return nil
}
