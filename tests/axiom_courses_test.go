package tests

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

type axiomCourse struct {
	Title     string
	Published bool
}

var coursesRunner = testRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag("courses"),
		),
	),
)

func (s *AxiomSuite) TestCourseCanBeCreated() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be created"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("create"),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		course := axiomCourse{
			Title:     "Go API Autotests",
			Published: false,
		}

		assert.Equal(
			cfg.T(),
			[]string{"axiom", "courses", "create"},
			cfg.Meta.Tags,
		)
		assert.Equal(cfg.T(), "Go API Autotests", course.Title)
		assert.False(cfg.T(), course.Published)
	})
}

func (s *AxiomSuite) TestCourseCanBePublished() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be published"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("publish"),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		course := axiomCourse{
			Title:     "Go API Autotests",
			Published: false,
		}

		assert.Equal(
			cfg.T(),
			[]string{"axiom", "courses", "publish"},
			cfg.Meta.Tags,
		)

		course.Published = true

		assert.True(cfg.T(), course.Published)
	})
}
