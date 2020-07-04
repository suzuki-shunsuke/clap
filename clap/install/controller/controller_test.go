package controller_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/clap/clap/install/controller"
	"github.com/suzuki-shunsuke/clap/clap/install/extract"
)

type mockExtractor struct {
	err error
}

func (m mockExtractor) Extract(params extract.ParamsExtract) error {
	return m.err
}

func TestController_Run(t *testing.T) {
	u, err := url.Parse("https://github.com/suzuki-shunsuke/cmdx/releases/download/v1.6.0/cmdx_1.6.0_linux_amd64.tar.gz")
	require.Nil(t, err)
	base, _, err := controller.New(controller.ParamsNew{})
	require.Nil(t, err)
	ctrl1 := base
	ctrl1.Extractor = mockExtractor{}

	data := []struct {
		title  string
		ctrl   controller.Controller
		params controller.ParamsRun
		isErr  bool
	}{
		{
			title: "normal",
			params: controller.ParamsRun{
				URL: u,
				Files: []controller.File{
					{
						Path:        "foo",
						InstallPath: "/tmp/foo",
					},
				},
			},
			ctrl: ctrl1,
		},
	}
	for _, d := range data {
		d := d
		ctx := context.Background()
		t.Run(d.title, func(t *testing.T) {
			err := d.ctrl.Run(ctx, d.params)
			if d.isErr {
				require.NotNil(t, err)
			}
			require.Nil(t, err)
		})
	}
}
