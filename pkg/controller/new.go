package controller

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/suzuki-shunsuke/clap/pkg/download"
	"github.com/suzuki-shunsuke/clap/pkg/fsys"
	"github.com/suzuki-shunsuke/clap/pkg/unarchiver"
)

type Controller struct {
	Downloader  Downloader
	Unarchiver  Unarchiver
	TempDir     TempDirCreator
	FileRemover FileRemover
	FileCreator FileCreator
	FileChecker FileChecker
	FileRenamer FileRenamer
	Mkdir       Mkdir
}

func New(params ParamsNew) (Controller, ParamsRun, error) {
	return Controller{
		Downloader:  download.New(download.ParamsNew{}),
		Unarchiver:  unarchiver.Unarchiver{},
		TempDir:     fsys.TempDir{},
		FileRemover: fsys.FileRemover{},
		FileCreator: fsys.FileCreator{},
		FileChecker: fsys.FileChecker{},
		FileRenamer: fsys.FileRenamer{},
		Mkdir:       fsys.Mkdir{},
	}, ParamsRun(params), nil
}

type ParamsNew struct {
	URL    *url.URL
	Files  []File
	Method string
	Header http.Header
}

type FileChecker interface {
	Stat(string) (os.FileInfo, error)
}

type Downloader interface {
	Run(ctx context.Context, params download.ParamsDownload) (*http.Response, error)
}

type TempDirCreator interface {
	Create() (string, error)
}

type FileRemover interface {
	RemoveAll(string) error
}

type Unarchiver interface {
	Unarchive(params unarchiver.ParamsUnarchive) error
}

type Permissioner interface {
	Chmod(string, os.FileMode) error
}

type FileCreator interface {
	Create(params fsys.ParamsCreateFile) error
}

type FileRenamer interface {
	Rename(src, dest string) error
}

type Mkdir interface {
	MkdirAll(path string, mode os.FileMode) error
}
