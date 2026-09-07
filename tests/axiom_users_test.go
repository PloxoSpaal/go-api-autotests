package tests

import (
	"testing"

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

func TestAxiomCreateUser(t *testing.T) {
	testCase := axiom.NewCase(
		axiom.WithCaseName("user can be created"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("create"),
		),
	)

	usersRunner.RunCase(t, testCase, func(cfg *axiom.Config) {
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

func TestAxiomDeactivateUser(t *testing.T) {
	testCase := axiom.NewCase(
		axiom.WithCaseName("user can be deactivated"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("deactivate"),
		),
	)

	usersRunner.RunCase(t, testCase, func(cfg *axiom.Config) {
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
