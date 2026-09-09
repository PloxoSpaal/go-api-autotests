package tests

import "github.com/Nikita-Filonov/axiom"

type axiomCoursesTools struct {
	cfg *axiom.Config

	Client *axiomCourseClient

	Environment string
	Service     string
	Operation   string
}

func (t *axiomCoursesTools) CourseData() axiomCourseData {
	return axiom.GetFixture[axiomCourseData](t.cfg, "course-data")
}

func (t *axiomCoursesTools) UnpublishedCourse() *axiomCourse {
	return axiom.GetFixture[*axiomCourse](t.cfg, "unpublished-course")
}

var coursesTools = axiom.NewToolset("course.tools", func(cfg *axiom.Config) *axiomCoursesTools {
	return &axiomCoursesTools{
		cfg:         cfg,
		Client:      axiom.MustResource[*axiomCourseClient](cfg.Runner, "course-client"),
		Environment: axiom.MustContextValue[string](&cfg.Context, "environment"),
		Service:     axiom.MustContextValue[string](&cfg.Context, "service"),
		Operation:   axiom.MustContextValue[string](&cfg.Context, "operation"),
	}
})
