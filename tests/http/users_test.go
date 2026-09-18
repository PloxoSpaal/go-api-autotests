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
		user := tools.Builder.UserCreate()
		request := user.HTTPRequest()

		response, err := tools.PublicUsersClient().Create(cfg.Context.Raw, request)
		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)
		require.NotEmpty(cfg.T(), response.Data.User.ID)

		assert.Equal(cfg.T(), request.Email, response.Data.User.Email)
		assert.Equal(cfg.T(), request.LastName, response.Data.User.LastName)
		assert.Equal(cfg.T(), request.FirstName, response.Data.User.FirstName)
		assert.Equal(cfg.T(), request.MiddleName, response.Data.User.MiddleName)

		cfg.T().Logf("created user with ID %s", response.Data.User.ID)
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
		request := update.HTTPRequest()

		response, err := tools.PrivateUsersClient().Update(
			cfg.Context.Raw,
			userFixture.Response.Data.User.ID,
			request,
		)
		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)

		assert.Equal(cfg.T(), userFixture.Response.Data.User.ID, response.Data.User.ID)
		assert.Equal(cfg.T(), *request.Email, response.Data.User.Email)
		assert.Equal(cfg.T(), *request.LastName, response.Data.User.LastName)
		assert.Equal(cfg.T(), *request.FirstName, response.Data.User.FirstName)
		assert.Equal(cfg.T(), *request.MiddleName, response.Data.User.MiddleName)

		cfg.T().Logf("updated user with ID %s", response.Data.User.ID)
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
