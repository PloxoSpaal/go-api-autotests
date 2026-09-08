package tests

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

var usersRunner = testRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaFeature("Users"),
			axiom.WithMetaSuite("User management"),
			axiom.WithMetaSubSuite("Lifecycle"),
			axiom.WithMetaTag("users"),
			axiom.WithMetaLabel("owner", "users-team"),
			axiom.WithMetaLabel("component", "users-service"),
		),
		axiom.WithRunnerContext(axiom.WithContextData("service", "users-service")),
		axiom.WithRunnerResource("user-client", userClientResource),
		axiom.WithRunnerFixture("user-data", userDataFixture),
		axiom.WithRunnerFixture("active-user", activeUserFixture),
	),
)

func (s *AxiomSuite) TestUserCanBeCreated() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("user can be created"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("Create user"),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
			axiom.WithMetaTags("create", "smoke"),
			axiom.WithMetaIssue("AXIOM-101"),
			axiom.WithMetaTestCase("USERS-001"),
		),
		axiom.WithCaseContext(
			axiom.WithContextRaw(s.T().Context()),
			axiom.WithContextData("operation", "create-user"),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		client := axiom.MustResource[*axiomUserClient](cfg.Runner, "user-client")
		data := axiom.GetFixture[axiomUserData](cfg, "user-data")

		var user *axiomUser

		cfg.Step("check test metadata", func() {
			assert.Equal(cfg.T(), "Learning platform", cfg.Meta.Epic)
			assert.Equal(cfg.T(), "Users", cfg.Meta.Feature)
			assert.Equal(cfg.T(), "Create user", cfg.Meta.Story)
			assert.Equal(cfg.T(), axiom.SeverityCritical, cfg.Meta.Severity)
			assert.Equal(cfg.T(), []string{"axiom", "users", "create", "smoke"}, cfg.Meta.Tags)
			assert.Equal(cfg.T(), []string{"AXIOM-101"}, cfg.Meta.Issues)
			assert.Equal(cfg.T(), []string{"USERS-001"}, cfg.Meta.TestCases)

			expectedLabels := map[string]string{
				"owner":     "users-team",
				"component": "users-service",
			}
			assert.Equal(cfg.T(), expectedLabels, cfg.Meta.Labels)
		})

		cfg.Step("check test context", func() {
			environment := axiom.MustContextValue[string](&cfg.Context, "environment")
			service := axiom.MustContextValue[string](&cfg.Context, "service")
			operation := axiom.MustContextValue[string](&cfg.Context, "operation")

			assert.Equal(cfg.T(), "local", environment)
			assert.Equal(cfg.T(), "users-service", service)
			assert.Equal(cfg.T(), "create-user", operation)
		})

		cfg.Step("create user", func() {
			user = client.Create(data)
		})

		cfg.Step("check created user", func() {
			assert.Equal(cfg.T(), "student@example.com", user.Email)
			assert.True(cfg.T(), user.Active)
		})
	})
}

func (s *AxiomSuite) TestUserCanBeDeactivated() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("user can be deactivated"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("Deactivate user"),
			axiom.WithMetaTags("deactivate", "regression"),
			axiom.WithMetaIssue("AXIOM-102"),
			axiom.WithMetaTestCase("USERS-002"),
		),
		axiom.WithCaseContext(
			axiom.WithContextRaw(s.T().Context()),
			axiom.WithContextData("operation", "deactivate-user"),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		client := axiom.MustResource[*axiomUserClient](cfg.Runner, "user-client")
		user := axiom.GetFixture[*axiomUser](cfg, "active-user")

		cfg.Step("check test metadata", func() {
			assert.Equal(cfg.T(), "Deactivate user", cfg.Meta.Story)
			assert.Equal(cfg.T(), axiom.SeverityNormal, cfg.Meta.Severity)
			assert.Equal(cfg.T(), []string{"axiom", "users", "deactivate", "regression"}, cfg.Meta.Tags)
			assert.Equal(cfg.T(), []string{"AXIOM-102"}, cfg.Meta.Issues)
			assert.Equal(cfg.T(), []string{"USERS-002"}, cfg.Meta.TestCases)
		})

		cfg.Step("check test context", func() {
			environment := axiom.MustContextValue[string](&cfg.Context, "environment")
			service := axiom.MustContextValue[string](&cfg.Context, "service")
			operation := axiom.MustContextValue[string](&cfg.Context, "operation")

			assert.Equal(cfg.T(), "local", environment)
			assert.Equal(cfg.T(), "users-service", service)
			assert.Equal(cfg.T(), "deactivate-user", operation)
		})

		cfg.Step("deactivate user", func() {
			client.Deactivate(user)
		})

		cfg.Step("check deactivated user", func() {
			assert.False(cfg.T(), user.Active)
		})
	})
}
