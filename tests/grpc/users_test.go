package grpc

import (
	"context"
	"testing"

	v1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/fake"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func (s *suite) TestCreateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("create user"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		request := user.GRPCRequest()

		connection, err := grpc.NewClient(
			"localhost:9000",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		require.NoError(cfg.T(), err)
		defer connection.Close()

		client := v1.NewUsersServiceClient(connection)
		response, err := client.CreateUser(context.Background(), request)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())
		require.NotEmpty(cfg.T(), response.GetUser().GetId())

		assert.Equal(cfg.T(), request.GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), request.GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), request.GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), request.GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("created user with ID %s", response.GetUser().GetId())
	})
}

func TestUpdateUser(t *testing.T) {
	userRequest := &v1.CreateUserRequest{
		Email:      gofakeit.Email(),
		Password:   gofakeit.Password(true, true, true, true, false, 12),
		LastName:   gofakeit.LastName(),
		FirstName:  gofakeit.FirstName(),
		MiddleName: gofakeit.FirstName(),
	}

	connection, err := grpc.NewClient(
		"localhost:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	defer connection.Close()

	usersClient := v1.NewUsersServiceClient(connection)
	authenticationClient := v1.NewAuthenticationServiceClient(connection)

	createdUser, err := usersClient.CreateUser(context.Background(), userRequest)

	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.NotNil(t, createdUser.GetUser())
	require.NotEmpty(t, createdUser.GetUser().GetId())

	loginRequest := &v1.LoginRequest{
		Email:    userRequest.GetEmail(),
		Password: userRequest.GetPassword(),
	}

	loginResponse, err := authenticationClient.Login(context.Background(), loginRequest)

	require.NoError(t, err)
	require.NotNil(t, loginResponse)
	require.NotNil(t, loginResponse.GetToken())
	require.NotEmpty(t, loginResponse.GetToken().GetAccessToken())

	request := &v1.UpdateUserRequest{
		Id:         createdUser.GetUser().GetId(),
		Email:      new(gofakeit.Email()),
		LastName:   new(gofakeit.LastName()),
		FirstName:  new(gofakeit.FirstName()),
		MiddleName: new(gofakeit.FirstName()),
	}

	authorizationMetadata := metadata.Pairs(
		"authorization",
		"Bearer "+loginResponse.GetToken().GetAccessToken(),
	)

	contextWithToken := metadata.NewOutgoingContext(context.Background(), authorizationMetadata)

	response, err := usersClient.UpdateUser(contextWithToken, request)

	require.NoError(t, err)
	require.NotNil(t, response)
	require.NotNil(t, response.GetUser())

	assert.Equal(t, createdUser.GetUser().GetId(), response.GetUser().GetId())
	assert.Equal(t, request.GetEmail(), response.GetUser().GetEmail())
	assert.Equal(t, request.GetLastName(), response.GetUser().GetLastName())
	assert.Equal(t, request.GetFirstName(), response.GetUser().GetFirstName())
	assert.Equal(t, request.GetMiddleName(), response.GetUser().GetMiddleName())

	t.Logf("updated user with ID %s", response.GetUser().GetId())
}

func TestGetUserMe(t *testing.T) {
	userRequest := &v1.CreateUserRequest{
		Email:      gofakeit.Email(),
		Password:   gofakeit.Password(true, true, true, true, false, 12),
		LastName:   gofakeit.LastName(),
		FirstName:  gofakeit.FirstName(),
		MiddleName: gofakeit.FirstName(),
	}

	connection, err := grpc.NewClient(
		"localhost:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	defer connection.Close()

	usersClient := v1.NewUsersServiceClient(connection)
	authenticationClient := v1.NewAuthenticationServiceClient(connection)

	createdUser, err := usersClient.CreateUser(context.Background(), userRequest)

	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.NotNil(t, createdUser.GetUser())
	require.NotEmpty(t, createdUser.GetUser().GetId())

	loginRequest := &v1.LoginRequest{
		Email:    userRequest.GetEmail(),
		Password: userRequest.GetPassword(),
	}

	loginResponse, err := authenticationClient.Login(context.Background(), loginRequest)

	require.NoError(t, err)
	require.NotNil(t, loginResponse)
	require.NotNil(t, loginResponse.GetToken())
	require.NotEmpty(t, loginResponse.GetToken().GetAccessToken())

	request := &v1.Empty{}

	authorizationMetadata := metadata.Pairs(
		"authorization",
		"Bearer "+loginResponse.GetToken().GetAccessToken(),
	)

	contextWithToken := metadata.NewOutgoingContext(context.Background(), authorizationMetadata)

	response, err := usersClient.GetMe(contextWithToken, request)

	require.NoError(t, err)
	require.NotNil(t, response)
	require.NotNil(t, response.GetUser())

	assert.Equal(t, createdUser.GetUser().GetId(), response.GetUser().GetId())
	assert.Equal(t, userRequest.GetEmail(), response.GetUser().GetEmail())
	assert.Equal(t, userRequest.GetLastName(), response.GetUser().GetLastName())
	assert.Equal(t, userRequest.GetFirstName(), response.GetUser().GetFirstName())
	assert.Equal(t, userRequest.GetMiddleName(), response.GetUser().GetMiddleName())

	t.Logf("received user with ID %s", response.GetUser().GetId())
}
