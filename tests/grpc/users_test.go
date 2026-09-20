package grpc

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
)

func (s *suite) TestCreateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-USERS-001"),
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
		axiom.WithCaseID("GRPC-USERS-002"),
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
		axiom.WithCaseID("GRPC-USERS-003"),
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
		tools.Assertion.NoError(err)
		tools.Assertion.GRPCGetUserResponse(response, userFixture.Response)
	}))
}
