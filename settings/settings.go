package settings

import (
	"github.com/DavidHoenisch/gwhich/models/config"
)

type Settings struct {

	// NOTE: may not like this...also may not be possible
	// defines whether natural search features are enabled.
	// this relies on somewhat good descriptions of what
	// what a command does
	NaturalSearch bool

	Apps []config.Application
}

func New() *Settings {

	return &Settings{
		Apps: nil,
	}
}
