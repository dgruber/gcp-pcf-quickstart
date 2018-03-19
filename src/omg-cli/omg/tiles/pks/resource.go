package pks

import (
	"log"
	"omg-cli/config"
)

var tile = config.Tile{
	config.PivnetMetadata{
		"pivotal-container-service",
		43085,
		75338,
		"593bf193838ec63cc9e754e4df8b32d4f5e55430402de65d38defc1bb4ff465a",
	},
	config.OpsManagerMetadata{
		"pivotal-container-service",
		"1.0.0-build.3",
	},
	&config.StemcellMetadata{
		config.PivnetMetadata{"stemcells",
			36314,
			67872,
			"6c966883018e34edc8c0c61a48b7aa07582571e39e37f9065ff58eff4f4b4423"},
		"light-bosh-stemcell-3468.21-google-kvm-ubuntu-trusty-go_agent",
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
