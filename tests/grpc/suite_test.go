package grpc

import (
	"testing"

	"github.com/Nikita-Filonov/axiom"
)

func TestSuite(t *testing.T) {
	testSuite := axiom.NewSuiteFactory(
		t,
		func() *suite { return new(suite) },
		axiom.WithSuiteConfigRunner(suiteRunner),
		axiom.WithSuiteConfigParallel(),
	)

	testSuite.Test(
		"TestCreateUser",
		(*suite).TestCreateUser,
		axiom.WithSuiteTestRunner(usersRunner),
	)
	testSuite.Test(
		"TestUpdateUser",
		(*suite).TestUpdateUser,
		axiom.WithSuiteTestRunner(usersRunner),
	)
	testSuite.Test(
		"TestGetUserMe",
		(*suite).TestGetUserMe,
		axiom.WithSuiteTestRunner(usersRunner),
	)

	testSuite.Test(
		"TestLogin",
		(*suite).TestLogin,
		axiom.WithSuiteTestRunner(authenticationRunner),
	)

	testSuite.Run()
}
