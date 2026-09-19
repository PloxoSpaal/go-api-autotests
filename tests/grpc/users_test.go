package grpc

import (
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
		response, err := tools.PublicUsersClient().Create(
			cfg.Context.Raw,
			create.GRPCRequest(),
		)

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCCreateUserResponse(response, create)
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
			update.GRPCRequest(userFixture.Response.GetUser().GetId()),
		)

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCUpdateUserResponse(response, update)
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
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())

		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetId(), response.GetUser().GetId())
		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), userFixture.Response.GetUser().GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("received current user with ID %s", response.GetUser().GetId())
	}))
}
