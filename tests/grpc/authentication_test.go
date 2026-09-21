package grpc

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
	"google.golang.org/grpc/codes"
)

func (s *suite) TestLogin() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-AUTHENTICATION-001"),
		axiom.WithCaseName("login with correct email and password"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryLogin),
			axiom.WithMetaSeverity(axiom.SeverityBlocker),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		userFixture := tools.User()

		request := &v1.LoginRequest{
			Email:    userFixture.Request.GetEmail(),
			Password: userFixture.Request.GetPassword(),
		}
		response, err := tools.AuthenticationClient().Login(cfg.Context.Raw, request)

		tools.Assertion.NoError(err)
		tools.Assertion.GRPCLoginResponse(response)
	}))
}

func (s *suite) TestLoginWithInvalidCredentials() {
	testCase := axiom.NewCase(
		axiom.WithCaseID("GRPC-AUTHENTICATION-002"),
		axiom.WithCaseName("login with incorrect email and password"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagNegative),
			axiom.WithMetaStory(metadata.StoryLogin),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, suiteToolset.Action(func(cfg *axiom.Config, tools *suiteTools) {
		request := &v1.LoginRequest{
			Email:    tools.Fake.Email(),
			Password: tools.Fake.Password(),
		}
		response, err := tools.AuthenticationClient().Login(cfg.Context.Raw, request)

		tools.Assertion.GRPCError(err, codes.Unauthenticated, "Invalid credentials")
		tools.Assertion.Nil(response, "gRPC response")
	}))
}
