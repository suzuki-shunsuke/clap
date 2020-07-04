package cli_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/clap/clap/install/cli"
	"github.com/suzuki-shunsuke/clap/clap/install/controller"
)

func TestParseInput(t *testing.T) {
	urlString := "http://example.com"
	u, err := url.Parse(urlString)
	require.Nil(t, err)
	data := []struct {
		title string
		input cli.Input
		isErr bool
		exp   controller.ParamsNew
	}{
		{
			title: "normal",
			input: cli.Input{
				URL: urlString,
				Files: []string{
					"foo:/tmp/foo",
				},
			},
			exp: controller.ParamsNew{
				URL: u,
				Files: []controller.File{
					{
						Path:        "foo",
						InstallPath: "/tmp/foo",
					},
				},
			},
		},
		{
			title: "invalid argument",
			input: cli.Input{
				URL: urlString,
				Files: []string{
					"foo",
				},
			},
			isErr: true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			paramsNew, err := cli.ParseInput(d.input)
			if d.isErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			require.Equal(t, d.exp, paramsNew)
		})
	}
}
