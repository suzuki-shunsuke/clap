package unarchiver

import (
	"github.com/mholt/archiver/v3"
)

type Unarchiver struct{}

type ParamsUnarchive struct {
	Source      string
	Destination string
}

func (Unarchiver) Unarchive(params ParamsUnarchive) error {
	return archiver.Unarchive(params.Source, params.Destination)
}
