package resources

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/PloxoSpaal/go-api-autotests/fake"
)

const FakeResourceKey = "fake"

func SetFakeResource(_ *axiom.Runner) (any, func(), error) {
	return fake.New(), nil, nil
}

func GetFakeResource(runner *axiom.Runner) *fake.Fake {
	return axiom.MustResource[*fake.Fake](runner, FakeResourceKey)
}
