package controller

import (
	"errors"
	"fmt"
	"net/http"
)

type ParamsValidateResp struct {
	Resp *http.Response
}

var ErrHTTPResponseStatusCode = errors.New("http response status code >= 300")

func (ctrl Controller) validateResp(params ParamsValidateResp) error {
	if params.Resp.StatusCode >= 300 { //nolint:gomnd
		return fmt.Errorf("%w (status code: %d)", ErrHTTPResponseStatusCode, params.Resp.StatusCode)
	}
	return nil
}
