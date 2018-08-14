package pks

import (
	"log"
	"omg-cli/config"
)

var tile = config.Tile{
	config.PivnetMetadata{
		"pivotal-container-service",
		156571,
		190794,
		"a06c8309936bf759decd76a984a9968e325a7cf58bac5a60bb45fabfd74bed5e",
	},
	config.OpsManagerMetadata{
		"pivotal-container-service",
		"1.1.4-build.5",
	},
	&config.StemcellMetadata{
		config.PivnetMetadata{"stemcells",
			151610,
			187238,
			"d8c6a1a2b955c56f796238e7c7a0c27d165896a18d981d28f5acb9d09d1fc869"},
		"light-bosh-stemcell-3586.27-google-kvm-ubuntu-trusty-go_agent",
	},
}

type Tile struct {
	Logger *log.Logger
}

func (*Tile) Definition(*config.EnvConfig) config.Tile {
	return tile
}

func (*Tile) BuiltIn() bool {
	return false
}

func (*Tile) NoConfig() bool {
	return false
}
