package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	"github.com/PloxoSpaal/go-api-autotests/transport/httptransport"
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

	s.RunCase(testCase, func(cfg *axiom.Config) {
		settings := resources.GetConfigResource(cfg.Runner)

		publicTransport := httptransport.NewPublic(cfg, settings.HTTP)
		usersClient := httpclients.NewUsersClient(publicTransport)

		generator := resources.GetFakeResource(cfg.Runner)
		builder := builders.New(generator)
		user := builder.UserCreate()
		request := user.HTTPRequest()

		response, err := usersClient.Create(cfg.Context.Raw, request)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)
		require.NotEmpty(cfg.T(), response.Data.User.ID)

		assert.Equal(cfg.T(), request.Email, response.Data.User.Email)
		assert.Equal(cfg.T(), request.LastName, response.Data.User.LastName)
		assert.Equal(cfg.T(), request.FirstName, response.Data.User.FirstName)
		assert.Equal(cfg.T(), request.MiddleName, response.Data.User.MiddleName)

		cfg.T().Logf("created user with ID %s", response.Data.User.ID)
	})
}

func (s *suite) TestUpdateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("update user"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryUpdateEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		settings := resources.GetConfigResource(cfg.Runner)

		publicTransport := httptransport.NewPublic(cfg, settings.HTTP)
		usersPublicClient := httpclients.NewUsersClient(publicTransport)
		authenticationClient := httpclients.NewAuthenticationClient(publicTransport)

		generator := resources.GetFakeResource(cfg.Runner)
		builder := builders.New(generator)
		user := builder.UserCreate()
		userRequest := user.HTTPRequest()

		createdUserResponse, err := usersPublicClient.Create(cfg.Context.Raw, userRequest)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, createdUserResponse.StatusCode)
		require.NotEmpty(cfg.T(), createdUserResponse.Data.User.ID)

		authRequest := models.LoginRequest{
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		loginHTTPResponse, err := authenticationClient.Login(cfg.Context.Raw, authRequest)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, loginHTTPResponse.StatusCode)
		require.NotEmpty(cfg.T(), loginHTTPResponse.Data.Token.AccessToken)

		privateTransport := httptransport.NewPrivate(cfg, settings.HTTP, loginHTTPResponse.Data.Token.AccessToken)
		usersPrivateClient := httpclients.NewUsersClient(privateTransport)

		update := builder.UserUpdate()
		request := update.HTTPRequest()

		response, err := usersPrivateClient.Update(cfg.Context.Raw, createdUserResponse.Data.User.ID, request)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)

		assert.Equal(cfg.T(), createdUserResponse.Data.User.ID, response.Data.User.ID)
		assert.Equal(cfg.T(), *request.Email, response.Data.User.Email)
		assert.Equal(cfg.T(), *request.LastName, response.Data.User.LastName)
		assert.Equal(cfg.T(), *request.FirstName, response.Data.User.FirstName)
		assert.Equal(cfg.T(), *request.MiddleName, response.Data.User.MiddleName)

		cfg.T().Logf("updated user with ID %s", response.Data.User.ID)
	})
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

	s.RunCase(testCase, func(cfg *axiom.Config) {
		settings := resources.GetConfigResource(cfg.Runner)

		publicTransport := httptransport.NewPublic(cfg, settings.HTTP)
		usersPublicClient := httpclients.NewUsersClient(publicTransport)
		authenticationClient := httpclients.NewAuthenticationClient(publicTransport)

		generator := resources.GetFakeResource(cfg.Runner)
		builder := builders.New(generator)
		user := builder.UserCreate()
		userRequest := user.HTTPRequest()

		createdUserResponse, err := usersPublicClient.Create(cfg.Context.Raw, userRequest)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, createdUserResponse.StatusCode)
		require.NotEmpty(cfg.T(), createdUserResponse.Data.User.ID)

		authRequest := models.LoginRequest{
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		loginHTTPResponse, err := authenticationClient.Login(cfg.Context.Raw, authRequest)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, loginHTTPResponse.StatusCode)
		require.NotEmpty(cfg.T(), loginHTTPResponse.Data.Token.AccessToken)

		privateTransport := httptransport.NewPrivate(cfg, settings.HTTP, loginHTTPResponse.Data.Token.AccessToken)
		usersPrivateClient := httpclients.NewUsersClient(privateTransport)

		response, err := usersPrivateClient.GetMe(cfg.Context.Raw)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)

		assert.Equal(cfg.T(), createdUserResponse.Data.User.ID, response.Data.User.ID)
		assert.Equal(cfg.T(), userRequest.Email, response.Data.User.Email)
		assert.Equal(cfg.T(), userRequest.LastName, response.Data.User.LastName)
		assert.Equal(cfg.T(), userRequest.FirstName, response.Data.User.FirstName)
		assert.Equal(cfg.T(), userRequest.MiddleName, response.Data.User.MiddleName)

		cfg.T().Logf("received user with ID %s", response.Data.User.ID)
	})
}
