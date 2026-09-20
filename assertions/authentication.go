package assertions

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPLoginResponse(actual models.LoginResponse) {
	message := "Check HTTP login response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.HTTPToken(actual.Token)
	})
}

func (a *Assertion) HTTPToken(actual models.Token) {
	message := "Check HTTP token"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.Equal(actual.TokenType, "bearer", "token type")
		a.NotEmpty(actual.AccessToken, "access token")
		a.NotEmpty(actual.RefreshToken, "refresh token")
	})
}

func (a *Assertion) GRPCLoginResponse(actual *v1.LoginResponse) {
	message := "Check gRPC login response"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "login response")
		a.GRPCToken(actual.GetToken())
	})
}

func (a *Assertion) GRPCToken(actual *v1.Token) {
	message := "Check gRPC token"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.NotNil(actual, "token")
		a.Equal(actual.GetTokenType(), "bearer", "token type")
		a.NotEmpty(actual.GetAccessToken(), "access token")
		a.NotEmpty(actual.GetRefreshToken(), "refresh token")
	})
}
