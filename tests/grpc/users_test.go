package grpc

import (
	"context"

	v1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	"github.com/PloxoSpaal/go-api-autotests/transport/grpctransport"
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

		connection, err := grpctransport.NewPublic(cfg, settings.GRPC)
		require.NoError(cfg.T(), err)
		defer connection.Close()

		usersClient := grpcclients.NewUsersClient(connection)

		generator := resources.GetFakeResource(cfg.Runner)
		builder := builders.New(generator)
		user := builder.UserCreate()
		request := user.GRPCRequest()

		response, err := usersClient.Create(cfg.Context.Raw, request)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())
		require.NotEmpty(cfg.T(), response.GetUser().GetId())

		assert.Equal(cfg.T(), request.GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), request.GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), request.GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), request.GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("created user with ID %s", response.GetUser().GetId())
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

		publicConnection, err := grpctransport.NewPublic(cfg, settings.GRPC)
		require.NoError(cfg.T(), err)
		defer publicConnection.Close()

		publicUsersClient := grpcclients.NewUsersClient(publicConnection)
		authenticationClient := grpcclients.NewAuthenticationClient(publicConnection)

		generator := resources.GetFakeResource(cfg.Runner)
		builder := builders.New(generator)
		user := builder.UserCreate()
		userRequest := user.GRPCRequest()

		createdUser, err := publicUsersClient.Create(cfg.Context.Raw, userRequest)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createdUser)
		require.NotNil(cfg.T(), createdUser.GetUser())
		require.NotEmpty(cfg.T(), createdUser.GetUser().GetId())

		loginResponse, err := authenticationClient.Login(
			cfg.Context.Raw,
			&v1.LoginRequest{
				Email:    user.Email,
				Password: user.Password,
			},
		)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), loginResponse)
		require.NotNil(cfg.T(), loginResponse.GetToken())
		require.NotEmpty(cfg.T(), loginResponse.GetToken().GetAccessToken())

		privateConnection, err := grpctransport.NewPrivate(
			cfg,
			settings.GRPC,
			loginResponse.GetToken().GetAccessToken(),
		)
		require.NoError(cfg.T(), err)
		defer privateConnection.Close()

		privateUsersClient := grpcclients.NewUsersClient(privateConnection)

		update := builder.UserUpdate()
		request := update.GRPCRequest(createdUser.GetUser().GetId())

		response, err := privateUsersClient.Update(cfg.Context.Raw, request)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())

		assert.Equal(cfg.T(), createdUser.GetUser().GetId(), response.GetUser().GetId())
		assert.Equal(cfg.T(), request.GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), request.GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), request.GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), request.GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("updated user with ID %s", response.GetUser().GetId())
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

		publicConnection, err := grpctransport.NewPublic(cfg, settings.GRPC)
		require.NoError(cfg.T(), err)
		defer publicConnection.Close()

		generator := resources.GetFakeResource(cfg.Runner)
		builder := builders.New(generator)
		user := builder.UserCreate()
		userRequest := user.GRPCRequest()

		publicUsersClient := grpcclients.NewUsersClient(publicConnection)
		publicAuthenticationClient := grpcclients.NewAuthenticationClient(publicConnection)

		createdUser, err := publicUsersClient.Create(cfg.Context.Raw, userRequest)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createdUser)
		require.NotNil(cfg.T(), createdUser.GetUser())
		require.NotEmpty(cfg.T(), createdUser.GetUser().GetId())

		loginRequest := &v1.LoginRequest{
			Email:    userRequest.GetEmail(),
			Password: userRequest.GetPassword(),
		}

		loginResponse, err := publicAuthenticationClient.Login(context.Background(), loginRequest)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), loginResponse)
		require.NotNil(cfg.T(), loginResponse.GetToken())
		require.NotEmpty(cfg.T(), loginResponse.GetToken().GetAccessToken())

		privateConnection, err := grpctransport.NewPrivate(
			cfg,
			settings.GRPC,
			loginResponse.GetToken().GetAccessToken(),
		)
		require.NoError(cfg.T(), err)
		defer privateConnection.Close()

		privateUsersClient := grpcclients.NewUsersClient(privateConnection)

		response, err := privateUsersClient.GetMe(context.Background())

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())

		assert.Equal(cfg.T(), createdUser.GetUser().GetId(), response.GetUser().GetId())
		assert.Equal(cfg.T(), userRequest.GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), userRequest.GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), userRequest.GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), userRequest.GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("received user with ID %s", response.GetUser().GetId())
	})
}
