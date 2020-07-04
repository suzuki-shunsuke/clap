package cli

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/suzuki-shunsuke/clap/clap/install/controller"
)

type Input struct {
	URL   string
	Files []string
}

var ErrTooFewArgument = errors.New("arguments is too few")
var ErrColonNotFound = errors.New("invalid argument. colon(:) is needed. the format should be <file path in archive>:<install path>")

func GetInput(cmd *cobra.Command, args []string) (Input, error) {
	if len(args) < 2 { //nolint:gomnd
		return Input{}, ErrTooFewArgument
	}
	return Input{
		URL:   args[0],
		Files: args[1:],
	}, nil
}

func ParseInput(input Input) (controller.ParamsNew, error) {
	files := make([]controller.File, len(input.Files))
	for i, a := range input.Files {
		idx := strings.Index(a, ":")
		if idx == -1 {
			return controller.ParamsNew{}, fmt.Errorf("%w (arg = %s)", ErrColonNotFound, a)
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
	return controller.ParamsNew{
		URL:   u,
		Files: files,
	}, nil
}
