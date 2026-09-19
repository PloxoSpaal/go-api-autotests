package grpc

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/metadata"
)

func (s *suite) TestLogin() {
	testCase := axiom.NewCase(
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
