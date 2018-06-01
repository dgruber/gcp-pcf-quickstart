package pks

import (
	"log"
	"omg-cli/config"
)

var tile = config.Tile{
	config.PivnetMetadata{
		"pivotal-container-service",
		104377,
		143263,
		"d285a81caafd048776b4f7ce6d7659e0784e1cccab0a3e94d0314118ceaca2af",
	},
	config.OpsManagerMetadata{
		"pivotal-container-service",
		"1.0.4-build.5",
	},
	&config.StemcellMetadata{
		config.PivnetMetadata{"stemcells",
			97389,
			135526,
			"e573776443dd7001f404687aa3b76a566dd008d07d2d5ec008894e1a255b0f6f"},
		"light-bosh-stemcell-3468.42-google-kvm-ubuntu-trusty-go_agent",
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
