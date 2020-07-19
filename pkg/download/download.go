package download

import (
	"context"
	"net/http"
	"net/url"
)

type Downloader struct {
	httpClient *http.Client
}

type ParamsNew struct {
	HTTPClient *http.Client
}

func New(params ParamsNew) Downloader {
	httpClient := params.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return Downloader{
		httpClient: httpClient,
	}
}

type ParamsDownload struct {
	URL *url.URL
}

func (dl Downloader) Run(ctx context.Context, params ParamsDownload) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, params.URL.String(), nil)
	if err != nil {
		return nil, err
	}
	return dl.httpClient.Do(req)
}
