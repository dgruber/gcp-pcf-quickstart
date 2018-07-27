package pks

import (
	"log"
	"omg-cli/config"
)

var tile = config.Tile{
	config.PivnetMetadata{
		"pivotal-container-service",
		139502,
		178623,
		"29fa8687e7336fc9f5144f03ecd4f2fa7cbabda7798d618e5492e05793e3fa90",
	},
	config.OpsManagerMetadata{
		"pivotal-container-service",
		"1.1.2-build.2",
	},
	&config.StemcellMetadata{
		config.PivnetMetadata{"stemcells",
			129480,
			161625,
			"8de79a7436ce7e7f772beeb9bb65ea4d7dedb7e190e443d8281b6aebe5e73526"},
		"light-bosh-stemcell-3586.24-google-kvm-ubuntu-trusty-go_agent",
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
