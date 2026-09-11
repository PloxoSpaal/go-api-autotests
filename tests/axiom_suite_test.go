package tests

import (
	"testing"

	"github.com/Nikita-Filonov/axiom"
)

func TestAxiomSuite(t *testing.T) {
	testSuite := axiom.NewSuiteFactory(
		t,
		func() *AxiomSuite { return new(AxiomSuite) },
		axiom.WithSuiteConfigRunner(testRunner),
	)

	testSuite.Test(
		"user can be created",
		(*AxiomSuite).TestUserCanBeCreated,
		axiom.WithSuiteTestRunner(usersRunner),
	)
	testSuite.Test(
		"user can be deactivated",
		(*AxiomSuite).TestUserCanBeDeactivated,
		axiom.WithSuiteTestRunner(usersRunner),
	)

	testSuite.Test(
		"course can be created",
		(*AxiomSuite).TestCourseCanBeCreated,
		axiom.WithSuiteTestRunner(coursesRunner),
	)
	testSuite.Test(
		"course can be published",
		(*AxiomSuite).TestCourseCanBePublished,
		axiom.WithSuiteTestRunner(coursesRunner),
	)
	testSuite.Test(
		"user can be created with different emails",
		(*AxiomSuite).TestUserCanBeCreatedWithDifferentEmails,
		axiom.WithSuiteTestRunner(usersRunner),
	)
	testSuite.Test(
		"course can be created with different titles",
		(*AxiomSuite).TestCourseCanBeCreatedWithDifferentTitles,
		axiom.WithSuiteTestRunner(coursesRunner),
	)

	testSuite.Run()
}
