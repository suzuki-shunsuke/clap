package cli

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/clap/pkg/controller"
)

func Test_parseInput(t *testing.T) {
	urlString := "http://example.com"
	u, err := url.Parse(urlString)
	require.Nil(t, err)
	data := []struct {
		title string
		input Input
		isErr bool
		exp   controller.ParamsNew
	}{
		{
			title: "normal",
			input: Input{
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
				Header: http.Header{},
			},
		},
		{
			title: "invalid argument",
			input: Input{
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
			paramsNew, err := parseInput(d.input)
			if d.isErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			require.Equal(t, d.exp, paramsNew)
		})
	}
}
