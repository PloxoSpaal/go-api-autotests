package grpcfixtures

import (
	"github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/resources"
	"github.com/stretchr/testify/require"
)

const UserFixtureKey = "grpc-user"

type UserFixture struct {
	Data     builders.UserCreate
	Request  *v1.CreateUserRequest
	Response *v1.GetUserResponse
}

func SetUserFixture(cfg *axiom.Config) (any, func(), error) {
	builder := resources.GetBuilderResource(cfg.Runner)
	usersClient := GetPublicUsersClientFixture(cfg)

	var fixture UserFixture

	cfg.Setup("Create fixture user", func() {
		fixture.Data = builder.UserCreate()
		fixture.Request = fixture.Data.GRPCRequest()

		var err error
		fixture.Response, err = usersClient.Create(cfg.Context.Raw, fixture.Request)
		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), fixture.Response)
		require.NotNil(cfg.T(), fixture.Response.GetUser())
		require.NotEmpty(cfg.T(), fixture.Response.GetUser().GetId())
	})

	return fixture, nil, nil
}

func GetUserFixture(cfg *axiom.Config) UserFixture {
	return axiom.GetFixture[UserFixture](cfg, UserFixtureKey)
}
