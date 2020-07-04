package fsys

import (
	"io/ioutil"
)

type TempDir struct{}

func (TempDir) Create() (string, error) {
	return ioutil.TempDir("", "clap")
}
