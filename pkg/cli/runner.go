package cli

import (
	"context"
	"io"
	"net/http"

	"github.com/suzuki-shunsuke/clap/pkg/constant"
	"github.com/suzuki-shunsuke/clap/pkg/controller"
	"github.com/urfave/cli/v2"
)

type Runner struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (runner Runner) Run(ctx context.Context, args ...string) error {
	app := cli.App{
		Name:    "clap",
		Usage:   "simple installer. https://github.com/suzuki-shunsuke/clap",
		Version: constant.Version,
		Commands: []*cli.Command{
			{
				Name:   "install",
				Usage:  "download a file and extract files from downloaded file and install them",
				Action: runner.action,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "request",
						Aliases: []string{"X"},
						Usage:   "HTTP Method",
						Value:   http.MethodGet,
					},
					&cli.StringSliceFlag{
						Name:    "header",
						Aliases: []string{"H"},
						Usage:   "HTTP Header",
					},
				},
			},
		},
	}

	return app.RunContext(ctx, args)
}

func (runner Runner) action(c *cli.Context) error {
	// URL <file path in archive>:<install path> ...
	input, err := getInput(c)
	if err != nil {
		return err
	}
	paramsNew, err := parseInput(input)
	if err != nil {
		return err
	}
	ctrl, params, err := controller.New(paramsNew)
	if err != nil {
		return err
	}
	return ctrl.Run(c.Context, params)
}
