package tests

import "github.com/Nikita-Filonov/axiom"

type axiomUsersTools struct {
	cfg *axiom.Config

	Client *axiomUserClient

	Environment string
	Service     string
	Operation   string
}

func (t *axiomUsersTools) UserData() axiomUserData {
	return axiom.GetFixture[axiomUserData](t.cfg, "user-data")
}

func (t *axiomUsersTools) ActiveUser() *axiomUser {
	return axiom.GetFixture[*axiomUser](t.cfg, "active-user")
}

var usersToolset = axiom.NewToolset("users.tools", func(cfg *axiom.Config) *axiomUsersTools {
	return &axiomUsersTools{
		cfg:         cfg,
		Client:      axiom.MustResource[*axiomUserClient](cfg.Runner, "user-client"),
		Environment: axiom.MustContextValue[string](&cfg.Context, "environment"),
		Service:     axiom.MustContextValue[string](&cfg.Context, "service"),
		Operation:   axiom.MustContextValue[string](&cfg.Context, "operation"),
	}
})
