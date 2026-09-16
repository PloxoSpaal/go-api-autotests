package resources

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/config"
)

const ConfigResourceKey = "config"

func SetConfigResource(_ *axiom.Runner) (any, func(), error) {
	settings, err := config.Load()
	return settings, nil, err
}

func GetConfigResource(runner *axiom.Runner) *config.Config {
	return axiom.MustResource[*config.Config](runner, ConfigResourceKey)
}
