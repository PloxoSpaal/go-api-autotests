package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/fake"
	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *suite) TestLoginUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("Login user"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		userRequest := user.HTTPRequest()

		var createdUser models.UserResponse

		client := resty.New().SetBaseURL("http://localhost:8000/api/v1")

		createUserResponse, err := client.
			R().
			SetBody(userRequest).
			SetResult(&createdUser).
			Post("/users")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createUserResponse)
		require.Equal(cfg.T(), http.StatusOK, createUserResponse.StatusCode())
		require.NotEmpty(cfg.T(), createdUser.User.ID)

		request := models.LoginRequest{
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		var result models.LoginResponse

		response, err := client.
			R().
			SetBody(request).
			SetResult(&result).
			Post("/authentication/login")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode())

		assert.Equal(cfg.T(), "bearer", result.Token.TokenType)
		assert.NotEmpty(cfg.T(), result.Token.AccessToken)
		assert.NotEmpty(cfg.T(), result.Token.RefreshToken)

		cfg.T().Logf("user %s successfully logged in", userRequest.Email)
	})
}
