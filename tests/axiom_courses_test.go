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
			axiom.WithMetaFeature("Courses"),
			axiom.WithMetaSuite("Course management"),
			axiom.WithMetaSubSuite("Lifecycle"),
			axiom.WithMetaTag("courses"),
			axiom.WithMetaLabel("owner", "courses-team"),
			axiom.WithMetaLabel("component", "courses-service"),
		),
		axiom.WithRunnerContext(
			axiom.WithContextData(
				"service", "courses-service",
			),
		),
	),
)

func (s *AxiomSuite) TestCourseCanBeCreated() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be created"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("Create course"),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
			axiom.WithMetaTags("create", "smoke"),
			axiom.WithMetaIssue("AXIOM-201"),
			axiom.WithMetaTestCase("COURSES-001"),
		),
		axiom.WithCaseContext(
			axiom.WithContextRaw(s.T().Context()),
			axiom.WithContextData("operation", "create-course"),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		var (
			name   string
			course axiomCourse
		)

		cfg.Setup("prepare course data", func() {
			name = "Go API Autotests"
		})

		defer cfg.Teardown("clear course data", func() {
			name = ""
			course = axiomCourse{}
		})

		cfg.Step("check test context", func() {
			environment := axiom.MustContextValue[string](&cfg.Context, "environment")
			service := axiom.MustContextValue[string](&cfg.Context, "service")
			operation := axiom.MustContextValue[string](&cfg.Context, "operation")

			assert.Equal(cfg.T(), "local", environment)
			assert.Equal(cfg.T(), "courses-service", service)
			assert.Equal(cfg.T(), "create-course", operation)
		})

		cfg.Step("check test metadata", func() {
			assert.Equal(cfg.T(), "Courses", cfg.Meta.Feature)
			assert.Equal(cfg.T(), "Create course", cfg.Meta.Story)
			assert.Equal(cfg.T(), axiom.SeverityCritical, cfg.Meta.Severity)
			assert.Equal(
				cfg.T(),
				[]string{"axiom", "courses", "create", "smoke"},
				cfg.Meta.Tags,
			)
			assert.Equal(cfg.T(), []string{"AXIOM-201"}, cfg.Meta.Issues)
			assert.Equal(cfg.T(), []string{"COURSES-001"}, cfg.Meta.TestCases)
			assert.Equal(
				cfg.T(),
				map[string]string{"owner": "courses-team", "component": "courses-service"},
				cfg.Meta.Labels,
			)
		})

		cfg.Step("create course", func() {
			course = axiomCourse{
				Title:     name,
				Published: false,
			}
		})

		cfg.Step("check created course", func() {
			assert.Equal(cfg.T(), "Go API Autotests", course.Title)
			assert.False(cfg.T(), course.Published)
		})
	})
}

func (s *AxiomSuite) TestCourseCanBePublished() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be published"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("Publish course"),
			axiom.WithMetaTags("publish", "regression"),
			axiom.WithMetaIssue("AXIOM-202"),
			axiom.WithMetaTestCase("COURSES-002"),
		),
		axiom.WithCaseContext(
			axiom.WithContextRaw(s.T().Context()),
			axiom.WithContextData("operation", "publish-course"),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		var course axiomCourse

		cfg.Setup("prepare unpublished course", func() {
			course = axiomCourse{
				Title:     "Go API Autotests",
				Published: false,
			}
		})

		defer cfg.Teardown("clear course data", func() {
			course = axiomCourse{}
		})

		cfg.Step("check test context", func() {
			environment := axiom.MustContextValue[string](&cfg.Context, "environment")
			service := axiom.MustContextValue[string](&cfg.Context, "service")
			operation := axiom.MustContextValue[string](&cfg.Context, "operation")

			assert.Equal(cfg.T(), "local", environment)
			assert.Equal(cfg.T(), "courses-service", service)
			assert.Equal(cfg.T(), "publish-course", operation)
		})

		cfg.Step("check test metadata", func() {
			assert.Equal(cfg.T(), "Publish course", cfg.Meta.Story)
			assert.Equal(cfg.T(), axiom.SeverityNormal, cfg.Meta.Severity)
			assert.Equal(
				cfg.T(),
				[]string{"axiom", "courses", "publish", "regression"},
				cfg.Meta.Tags,
			)
			assert.Equal(cfg.T(), []string{"AXIOM-202"}, cfg.Meta.Issues)
			assert.Equal(cfg.T(), []string{"COURSES-002"}, cfg.Meta.TestCases)
		})

		cfg.Step("publish course", func() {
			course.Published = true
		})

		cfg.Step("check published course", func() {
			assert.True(cfg.T(), course.Published)
		})
	})
}
