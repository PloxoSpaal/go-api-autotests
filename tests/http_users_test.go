package tests

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type createUserRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

type user struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

type userResponse struct {
	User user `json:"user"`
}

type updateUserRequest struct {
	Email      string `json:"email"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

func TestHttpCreateUser(t *testing.T) {
	request := createUserRequest{
		Email:      gofakeit.Email(),
		Password:   gofakeit.Password(true, true, true, true, false, 12),
		LastName:   gofakeit.LastName(),
		FirstName:  gofakeit.FirstName(),
		MiddleName: gofakeit.MiddleName(),
	}

	client := resty.New().SetBaseURL("http://localhost:8000/api/v1")

	var result userResponse

	response, err := client.R().SetBody(request).SetResult(&result).Post("/users")

	require.NoError(t, err)
	require.NotNil(t, response)
	require.Equal(t, http.StatusOK, response.StatusCode())
	require.NotEmpty(t, result.User.Id)

	assert.Equal(t, request.Email, result.User.Email)
	assert.Equal(t, request.LastName, result.User.LastName)
	assert.Equal(t, request.FirstName, result.User.FirstName)
	assert.Equal(t, request.MiddleName, result.User.MiddleName)

	t.Logf("created user with ID %s", result.User.Id)
}

func TestHttpUpdateUser(t *testing.T) {
	userRequest := createUserRequest{
		Email:      gofakeit.Email(),
		Password:   gofakeit.Password(true, true, true, true, false, 12),
		LastName:   gofakeit.LastName(),
		FirstName:  gofakeit.FirstName(),
		MiddleName: gofakeit.FirstName(),
	}

	client := resty.New().
		SetBaseURL("http://localhost:8000/api/v1")

	var createdUser userResponse

	createUserResponse, err := client.
		R().
		SetBody(userRequest).
		SetResult(&createdUser).
		Post("/users")

	require.NoError(t, err)
	require.NotNil(t, createUserResponse)
	require.Equal(
		t,
		http.StatusOK,
		createUserResponse.StatusCode(),
	)
	require.NotEmpty(t, createdUser.User.Id)

	authRequest := loginRequest{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	var loginResult loginResponse

	loginHTTPResponse, err := client.
		R().
		SetBody(authRequest).
		SetResult(&loginResult).
		Post("/authentication/login")

	require.NoError(t, err)
	require.NotNil(t, loginHTTPResponse)
	require.Equal(t, http.StatusOK, loginHTTPResponse.StatusCode())
	require.NotEmpty(t, loginResult.Token.AccessToken)

	request := updateUserRequest{
		Email:      gofakeit.Email(),
		LastName:   gofakeit.LastName(),
		FirstName:  gofakeit.FirstName(),
		MiddleName: gofakeit.FirstName(),
	}

	updateUserURL := "/users/" + createdUser.User.Id

	var result userResponse

	response, err := client.
		R().
		SetHeader(
			"Authorization",
			"Bearer "+loginResult.Token.AccessToken,
		).
		SetBody(request).
		SetResult(&result).
		Patch(updateUserURL)

	require.NoError(t, err)
	require.NotNil(t, response)
	require.Equal(t, http.StatusOK, response.StatusCode())

	assert.Equal(t, createdUser.User.Id, result.User.Id)
	assert.Equal(t, request.Email, result.User.Email)
	assert.Equal(t, request.LastName, result.User.LastName)
	assert.Equal(t, request.FirstName, result.User.FirstName)
	assert.Equal(t, request.MiddleName, result.User.MiddleName)

	t.Logf("updated user with ID %s", result.User.Id)
}
