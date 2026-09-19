package assertions

import (
	v1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

func (a *Assertion) HTTPLoginResponse(actual models.LoginResponse) {
	a.cfg.Step("Check HTTP login response", func() {
		a.HTTPToken(actual.Token)
	})
}

func (a *Assertion) HTTPToken(actual models.Token) {
	a.cfg.Step("Check HTTP token", func() {
		a.Equal(actual.TokenType, "bearer", "token type")
		a.NotEmpty(actual.AccessToken, "access token")
		a.NotEmpty(actual.RefreshToken, "refresh token")
	})
}

func (a *Assertion) GRPCLoginResponse(actual *v1.LoginResponse) {
	a.cfg.Step("Check gRPC login response", func() {
		a.NotNil(actual, "login response")
		a.NotNil(actual.GetToken(), "token")
		a.GRPCToken(actual.GetToken())
	})
}

func (a *Assertion) GRPCToken(actual *v1.Token) {
	a.cfg.Step("Check gRPC token", func() {
		a.Equal(actual.GetTokenType(), "bearer", "token type")
		a.NotEmpty(actual.GetAccessToken(), "access token")
		a.NotEmpty(actual.GetRefreshToken(), "refresh token")
	})
}
