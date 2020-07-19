package extract

import (
	"github.com/mholt/archiver/v3"
)

type Extractor struct{}

type ParamsExtract struct {
	Source      string
	Target      string
	Destination string
}

func (Extractor) Extract(params ParamsExtract) error {
	return archiver.Extract(params.Source, params.Target, params.Destination)
}
