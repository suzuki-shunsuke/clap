package controller

import (
	"context"
	"fmt"
	"net/url"

	"github.com/suzuki-shunsuke/clap/pkg/download"
	"github.com/suzuki-shunsuke/clap/pkg/fsys"
)

type ParamsDownload struct {
	URL    *url.URL
	Source string
	Dir    string
	Method string
}

func (ctrl Controller) Download(ctx context.Context, params ParamsDownload) error {
	resp, err := ctrl.Downloader.Run(ctx, download.ParamsDownload{
		URL:    params.URL,
		Method: params.Method,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := ctrl.validateResp(ParamsValidateResp{
		Resp: resp,
	}); err != nil {
		return err
	}
	if err := ctrl.FileCreator.Create(fsys.ParamsCreateFile{
		Reader: resp.Body,
		Source: params.Source,
	}); err != nil {
		return fmt.Errorf("failed to download a file to %s: %w", params.Source, err)
	}
	return nil
}
