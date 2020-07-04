package controller

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/suzuki-shunsuke/clap/clap/install/download"
	"github.com/suzuki-shunsuke/clap/clap/install/extract"
	"github.com/suzuki-shunsuke/clap/clap/install/fsys"
)

type Controller struct {
	Downloader  Downloader
	Extractor   Extractor
	TempDir     TempDirCreator
	FileRemover FileRemover
	FileCreator FileCreator
	FileChecker FileChecker
}

func New(params ParamsNew) (Controller, ParamsRun, error) {
	return Controller{
		Downloader:  download.New(download.ParamsNew{}),
		Extractor:   extract.Extractor{},
		TempDir:     fsys.TempDir{},
		FileRemover: fsys.FileRemover{},
		FileCreator: fsys.FileCreator{},
		FileChecker: fsys.FileChecker{},
	}, ParamsRun(params), nil
}

type ParamsNew struct {
	URL   *url.URL
	Files []File
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

type Extractor interface {
	Extract(params extract.ParamsExtract) error
}

type Permissioner interface {
	Chmod(string, os.FileMode) error
}

type FileCreator interface {
	Create(params fsys.ParamsCreateFile) error
}
