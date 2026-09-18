package grpc

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/clients/grpcclients"
	"github.com/PloxoSpaal/go-api-autotests/fixtures/grpcfixtures"
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
		builder := resources.GetBuilderResource(cfg.Runner)
		usersClient := grpcfixtures.GetPublicUsersClientFixture(cfg)

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
		builder := resources.GetBuilderResource(cfg.Runner)

		userFixture := grpcfixtures.GetUserFixture(cfg)
		sessionFixture := grpcfixtures.GetSessionFixture(cfg)

		privateConnection, err := grpctransport.NewPrivate(
			cfg,
			settings.GRPC,
			sessionFixture.Response.GetToken().GetAccessToken(),
		)
		require.NoError(cfg.T(), err)
		defer privateConnection.Close()

		privateUsersClient := grpcclients.NewUsersClient(privateConnection)
		update := builder.UserUpdate()
		request := update.GRPCRequest(userFixture.Response.GetUser().GetId())

		response, err := privateUsersClient.Update(cfg.Context.Raw, request)
		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())

		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetId(), response.GetUser().GetId())
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

		userFixture := grpcfixtures.GetUserFixture(cfg)
		sessionFixture := grpcfixtures.GetSessionFixture(cfg)

		privateConnection, err := grpctransport.NewPrivate(
			cfg,
			settings.GRPC,
			sessionFixture.Response.GetToken().GetAccessToken(),
		)
		require.NoError(cfg.T(), err)
		defer privateConnection.Close()

		privateUsersClient := grpcclients.NewUsersClient(privateConnection)
		response, err := privateUsersClient.GetMe(cfg.Context.Raw)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())

		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetId(), response.GetUser().GetId())
		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("received current user with ID %s", response.GetUser().GetId())
	})
}
