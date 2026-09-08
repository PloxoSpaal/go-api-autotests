package tests

import "github.com/Nikita-Filonov/axiom"

type axiomUserData struct {
	Email string
}

type axiomUser struct {
	Email  string
	Active bool
}

type axiomUserClient struct{}

func (c *axiomUserClient) Create(data axiomUserData) *axiomUser {
	return &axiomUser{Email: data.Email, Active: true}
}

func (c *axiomUserClient) Deactivate(user *axiomUser) {
	user.Active = false
}

func (c *axiomUserClient) Delete(user *axiomUser) {
	*user = axiomUser{}
}

func userClientResource(_ *axiom.Runner) (any, func(), error) {
	return &axiomUserClient{}, nil, nil
}

func userDataFixture(cfg *axiom.Config) (any, func(), error) {
	var data axiomUserData

	cfg.Setup("prepare fixture user data", func() {
		data = axiomUserData{Email: "student@example.com"}
	})

	return data, nil, nil
}

func activeUserFixture(cfg *axiom.Config) (any, func(), error) {
	client := axiom.MustResource[*axiomUserClient](cfg.Runner, "user-client")
	data := axiom.GetFixture[axiomUserData](cfg, "user-data")

	var user *axiomUser

	cfg.Setup("create fixture active user", func() {
		user = client.Create(data)
	})

	cleanup := func() {
		cfg.Teardown("delete fixture active user", func() {
			client.Delete(user)
		})
	}

	return user, cleanup, nil
}
