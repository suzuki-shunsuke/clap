package controller_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/clap/pkg/controller"
	"github.com/suzuki-shunsuke/clap/pkg/unarchiver"
)

type mockUnarchiver struct {
	err error
}

type mockRenamer struct {
	err error
}

func (m mockUnarchiver) Unarchive(params unarchiver.ParamsUnarchive) error {
	return m.err
}

func (m mockRenamer) Rename(src, dest string) error {
	return m.err
}

func TestController_Run(t *testing.T) {
	u, err := url.Parse("https://github.com/suzuki-shunsuke/cmdx/releases/download/v1.6.0/cmdx_1.6.0_linux_amd64.tar.gz")
	require.Nil(t, err)
	base, _, err := controller.New(controller.ParamsNew{})
	require.Nil(t, err)
	ctrl1 := base
	ctrl1.Unarchiver = mockUnarchiver{}
	ctrl1.FileRenamer = mockRenamer{}

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
