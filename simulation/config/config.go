package config

import (
	"github.com/iotaledger/hive.go/configuration"
	"time"
)

// Config keys
const (
	numberNodes   = "NumberNodes"
	duration      = "Duration"
	arrowLifetime = "ArrowLifetime"
	vEnabled      = "VisualEnabled"
	dropOnUpdate  = "DropOnUpdate"
)

func NumberNodes(config *configuration.Configuration) int {
	return config.Int(numberNodes)
}

func Duration(config *configuration.Configuration) time.Duration {
	return time.Duration(config.Int(duration)) * time.Second
}

func ArrowLifetime(config *configuration.Configuration) time.Duration {
	return time.Duration(config.Int(arrowLifetime)) * time.Second
}

func DropOnUpdate(config *configuration.Configuration) bool {
	return config.Bool(dropOnUpdate)
}

func VisEnabled(config *configuration.Configuration) bool {
	return config.Bool(vEnabled)
}
