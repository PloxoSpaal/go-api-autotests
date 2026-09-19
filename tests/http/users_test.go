package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *suite) TestCreateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("create user"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryCreateEntity),
			axiom.WithMetaSeverity(axiom.SeverityBlocker),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		create := tools.Builder.UserCreate()
		response, err := tools.PublicUsersClient().Create(cfg.Context.Raw, create.HTTPRequest())

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPCreateUserResponse(response.Data, create)
	}))
}

func (s *suite) TestUpdateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("update user"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryUpdateEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		userFixture := tools.User()
		update := tools.Builder.UserUpdate()
		response, err := tools.PrivateUsersClient().Update(
			cfg.Context.Raw,
			userFixture.Response.Data.User.ID,
			update.HTTPRequest(),
		)

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPUpdateUserResponse(response.Data, update)
	}))
}

func (s *suite) TestGetUserMe() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("get user me"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryGetEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		userFixture := tools.User()

		response, err := tools.PrivateUsersClient().GetMe(cfg.Context.Raw)
		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)

		assert.Equal(cfg.T(), userFixture.Response.Data.User.ID, response.Data.User.ID)
		assert.Equal(cfg.T(), userFixture.Response.Data.User.Email, response.Data.User.Email)
		assert.Equal(cfg.T(), userFixture.Response.Data.User.LastName, response.Data.User.LastName)
		assert.Equal(cfg.T(), userFixture.Response.Data.User.FirstName, response.Data.User.FirstName)
		assert.Equal(cfg.T(), userFixture.Response.Data.User.MiddleName, response.Data.User.MiddleName)

		cfg.T().Logf("received current user with ID %s", response.Data.User.ID)
	}))
}
