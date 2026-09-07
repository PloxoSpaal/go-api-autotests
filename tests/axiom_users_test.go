package tests

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

type axiomUser struct {
	Email  string
	Active bool
}

var usersRunner = testRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag("users"),
		),
	),
)

func (s *AxiomSuite) TestUserCanBeCreated() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("user can be created"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("create"),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		user := axiomUser{
			Email:  "student@example.com",
			Active: true,
		}

		assert.Equal(
			cfg.T(),
			[]string{"axiom", "users", "create"},
			cfg.Meta.Tags,
		)
		assert.Equal(cfg.T(), "student@example.com", user.Email)
		assert.True(cfg.T(), user.Active)
	})
}

func (s *AxiomSuite) TestUserCanBeDeactivated() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("user can be deactivated"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("deactivate"),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		user := axiomUser{
			Email:  "student@example.com",
			Active: true,
		}

		user.Active = false

		assert.Equal(
			cfg.T(),
			[]string{"axiom", "users", "deactivate"},
			cfg.Meta.Tags,
		)
		assert.False(cfg.T(), user.Active)
	})
}
