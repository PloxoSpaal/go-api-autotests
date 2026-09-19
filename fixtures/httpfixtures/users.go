package httpfixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/builders"
	"github.com/PloxoSpaal/go-api-autotests/clients/httpclients"
	"github.com/PloxoSpaal/go-api-autotests/fixtures"
	"github.com/PloxoSpaal/go-api-autotests/models"
	"github.com/PloxoSpaal/go-api-autotests/resources"
)

const UserFixtureKey = "http-user"

type UserFixture struct {
	Data     builders.UserCreate
	Request  models.CreateUserRequest
	Response httpclients.Response[models.UserResponse]
}

func SetUserFixture(cfg *axiom.Config) (any, func(), error) {
	builder := resources.GetBuilderResource(cfg.Runner)
	assertion := fixtures.GetAssertionFixture(cfg)
	usersClient := GetPublicUsersClientFixture(cfg)

	var fixture UserFixture

	cfg.Setup("Create fixture user", func() {
		fixture.Data = builder.UserCreate()
		fixture.Request = fixture.Data.HTTPRequest()

		var err error
		fixture.Response, err = usersClient.Create(cfg.Context.Raw, fixture.Request)
		assertion.HTTPOK(fixture.Response.StatusCode, err)
		assertion.HTTPCreateUserResponse(fixture.Response.Data, fixture.Data)
	})

	return fixture, nil, nil
}

func GetUserFixture(cfg *axiom.Config) UserFixture {
	return axiom.GetFixture[UserFixture](cfg, UserFixtureKey)
}
