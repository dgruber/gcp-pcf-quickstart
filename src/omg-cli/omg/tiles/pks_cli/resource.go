package pks_cli

import (
	"log"
	"omg-cli/config"
)

var tile = config.Tile{
	config.PivnetMetadata{
		"pivotal-container-service",
		43085,
		75337,
		"51a06fe8655589fedb3d8d3f676c96e0069b17c004174a2d5782bce4f6d2b2d3",
	},
	config.OpsManagerMetadata{
		"pks-linux-amd64",
		"1.0.0-build.3",
	},
	nil,
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
	return true
}
