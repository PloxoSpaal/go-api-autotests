package http

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
		(*suite).TestLoginUser,
		axiom.WithSuiteTestRunner(authenticationRunner),
	)

	testSuite.Test(
		"TestLoginWithInvalidCredentials",
		(*suite).TestLoginWithInvalidCredentials,
		axiom.WithSuiteTestRunner(authenticationRunner),
	)

	testSuite.Test(
		"TestDeleteFile",
		(*suite).TestDeleteFile,
		axiom.WithSuiteTestRunner(filesRunner),
	)

	testSuite.Test(
		"TestCreateFile",
		(*suite).TestCreateFile,
		axiom.WithSuiteTestRunner(filesRunner),
	)

	testSuite.Test(
		"TestGetFile",
		(*suite).TestGetFile,
		axiom.WithSuiteTestRunner(filesRunner),
	)

	testSuite.Run()
}
