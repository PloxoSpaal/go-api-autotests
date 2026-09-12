package grpc

import (
	"context"

	v1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *suite) TestLogin() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("login with correct email and password"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		createUserRequest := user.GRPCRequest()

		connection, err := grpc.NewClient(
			"localhost:9000",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		require.NoError(cfg.T(), err)
		defer connection.Close()

		usersClient := v1.NewUsersServiceClient(connection)
		createdUser, err := usersClient.CreateUser(context.Background(), createUserRequest)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createdUser)
		require.NotNil(cfg.T(), createdUser.GetUser())
		require.NotEmpty(cfg.T(), createdUser.GetUser().GetId())

		request := &v1.LoginRequest{
			Email:    user.Email,
			Password: user.Password,
		}

		client := v1.NewAuthenticationServiceClient(connection)
		response, err := client.Login(context.Background(), request)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetToken())

		assert.Equal(cfg.T(), "bearer", response.GetToken().GetTokenType())
		assert.NotEmpty(cfg.T(), response.GetToken().GetAccessToken())
		assert.NotEmpty(cfg.T(), response.GetToken().GetRefreshToken())

		cfg.T().Logf("user %s successfully logged in", user.Email)
	})
}
