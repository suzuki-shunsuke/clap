package cli

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/suzuki-shunsuke/clap/pkg/controller"
	"github.com/urfave/cli/v2"
)

type Input struct {
	URL    string
	Files  []string
	Method string
	Header []string
}

var (
	errTooFewArgument = errors.New("arguments is too few")
	errColonNotFound  = errors.New("invalid argument. colon(:) is needed. the format should be <file path in archive>:<install path>")
	errInvalidHeader  = errors.New("invalid header. the parameter format should be '<key>:<value>'")
)

func getInput(c *cli.Context) (Input, error) {
	if c.Args().Len() < 2 { //nolint:gomnd
		return Input{}, errTooFewArgument
	}
	return Input{
		URL:    c.Args().First(),
		Files:  c.Args().Slice()[1:],
		Method: c.String("request"),
		Header: c.StringSlice("header"),
	}, nil
}

func parseInput(input Input) (controller.ParamsNew, error) {
	files := make([]controller.File, len(input.Files))
	for i, a := range input.Files {
		idx := strings.Index(a, ":")
		if idx == -1 {
			return controller.ParamsNew{}, fmt.Errorf("%w (arg = %s)", errColonNotFound, a)
		}
		files[i] = controller.File{
			Path:        a[:idx],
			InstallPath: a[idx+1:],
		}
	}
	u, err := url.Parse(input.URL)
	if err != nil {
		return controller.ParamsNew{}, err
	}
	header := http.Header{}
	for _, s := range input.Header {
		i := strings.Index(s, ":")
		if i < 1 {
			return controller.ParamsNew{}, errInvalidHeader
		}
		header.Add(s[:i], s[i+1:])
	}
	return controller.ParamsNew{
		URL:    u,
		Files:  files,
		Method: input.Method,
		Header: header,
	}, nil
}
