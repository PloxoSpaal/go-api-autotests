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

func (s *suite) TestCreateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("create user"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		request := user.HTTPRequest()

		var result models.UserResponse

		client := resty.New().SetBaseURL("http://localhost:8000/api/v1")

		response, err := client.R().SetBody(request).SetResult(&result).Post("/users")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode())
		require.NotEmpty(cfg.T(), result.User.ID)

		assert.Equal(cfg.T(), request.Email, result.User.Email)
		assert.Equal(cfg.T(), request.LastName, result.User.LastName)
		assert.Equal(cfg.T(), request.FirstName, result.User.FirstName)
		assert.Equal(cfg.T(), request.MiddleName, result.User.MiddleName)

		cfg.T().Logf("created user with ID %s", result.User.ID)
	})
}

func (s *suite) TestUpdateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("update user"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		update := builder.UserUpdate()
		userRequest := user.HTTPRequest()
		client := resty.New().SetBaseURL("http://localhost:8000/api/v1")

		var createdUser models.UserResponse

		createUserResponse, err := client.
			R().
			SetBody(userRequest).
			SetResult(&createdUser).
			Post("/users")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createUserResponse)
		require.Equal(
			cfg.T(),
			http.StatusOK,
			createUserResponse.StatusCode(),
		)
		require.NotEmpty(cfg.T(), createdUser.User.ID)

		authRequest := models.LoginRequest{
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		var loginResult models.LoginResponse

		loginHTTPResponse, err := client.
			R().
			SetBody(authRequest).
			SetResult(&loginResult).
			Post("/authentication/login")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), loginHTTPResponse)
		require.Equal(cfg.T(), http.StatusOK, loginHTTPResponse.StatusCode())
		require.NotEmpty(cfg.T(), loginResult.Token.AccessToken)

		request := update.HTTPRequest()

		updateUserURL := "/users/" + createdUser.User.ID

		var result models.UserResponse

		response, err := client.
			R().
			SetHeader(
				"Authorization",
				"Bearer "+loginResult.Token.AccessToken,
			).
			SetBody(request).
			SetResult(&result).
			Patch(updateUserURL)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode())

		assert.Equal(cfg.T(), createdUser.User.ID, result.User.ID)
		assert.Equal(cfg.T(), *request.Email, result.User.Email)
		assert.Equal(cfg.T(), *request.LastName, result.User.LastName)
		assert.Equal(cfg.T(), *request.FirstName, result.User.FirstName)
		assert.Equal(cfg.T(), *request.MiddleName, result.User.MiddleName)

		cfg.T().Logf("updated user with ID %s", result.User.ID)
	})
}

func (s *suite) TestGetUserMe() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("get user me"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		userRequest := user.HTTPRequest()
		client := resty.New().
			SetBaseURL("http://localhost:8000/api/v1")

		var createdUser models.UserResponse

		createUserResponse, err := client.
			R().
			SetBody(userRequest).
			SetResult(&createdUser).
			Post("/users")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createUserResponse)
		require.Equal(
			cfg.T(),
			http.StatusOK,
			createUserResponse.StatusCode(),
		)
		require.NotEmpty(cfg.T(), createdUser.User.ID)

		authRequest := models.LoginRequest{
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		var loginResult models.LoginResponse

		loginHTTPResponse, err := client.
			R().
			SetBody(authRequest).
			SetResult(&loginResult).
			Post("/authentication/login")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), loginHTTPResponse)
		require.Equal(cfg.T(), http.StatusOK, loginHTTPResponse.StatusCode())
		require.NotEmpty(cfg.T(), loginResult.Token.AccessToken)

		var result models.UserResponse

		response, err := client.
			R().
			SetHeader(
				"Authorization",
				"Bearer "+loginResult.Token.AccessToken,
			).
			SetResult(&result).
			Get("/users/me")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode())

		assert.Equal(cfg.T(), createdUser.User.ID, result.User.ID)
		assert.Equal(cfg.T(), userRequest.Email, result.User.Email)
		assert.Equal(cfg.T(), userRequest.LastName, result.User.LastName)
		assert.Equal(cfg.T(), userRequest.FirstName, result.User.FirstName)
		assert.Equal(cfg.T(), userRequest.MiddleName, result.User.MiddleName)

		cfg.T().Logf("received user with ID %s", result.User.ID)
	})
}
