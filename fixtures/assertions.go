package fixtures

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/assertions"
)

const AssertionFixtureKey = "assertion"

func SetAssertionFixture(cfg *axiom.Config) (any, func(), error) {
	return assertions.New(cfg), nil, nil
}

func GetAssertionFixture(cfg *axiom.Config) *assertions.Assertion {
	return axiom.GetFixture[*assertions.Assertion](cfg, AssertionFixtureKey)
}
