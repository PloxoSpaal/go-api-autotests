package assertions

import "github.com/PloxoSpaal/go-api-autotests/models"

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
