package http

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
)

func (s *suite) TestCreateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("HTTP-USERS-001"),
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
		axiom.WithCaseID("HTTP-USERS-002"),
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
		axiom.WithCaseID("HTTP-USERS-003"),
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

		tools.Assertion.HTTPOK(response.StatusCode, err)
		tools.Assertion.HTTPGetUserResponse(response.Data, userFixture.Response.Data)
	}))
}
