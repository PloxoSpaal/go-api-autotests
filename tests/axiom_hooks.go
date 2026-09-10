package tests

import (
	"fmt"

	"github.com/Nikita-Filonov/axiom"
)

func logBeforeAll(runner *axiom.Runner) {
	name := runner.Meta.Feature
	if name == "" {
		name = runner.Meta.ParentSuite
	}

	fmt.Printf("[BeforeAll] runner started: %s\n", name)
}

func logAfterAll(runner *axiom.Runner) {
	name := runner.Meta.Feature
	if name == "" {
		name = runner.Meta.ParentSuite
	}

	fmt.Printf("[AfterAll] runner finished: %s\n", name)
}

func logBeforeTest(cfg *axiom.Config) {
	cfg.T().Logf("[BeforeTest] test started: %s", cfg.Case.Name)
}

func logAfterTest(cfg *axiom.Config) {
	cfg.T().Logf("[AfterTest] test finished: %s", cfg.Case.Name)
}

func logBeforeStep(cfg *axiom.Config, name string) {
	cfg.T().Logf("[BeforeStep] step started: %s", name)
}

func logAfterStep(cfg *axiom.Config, name string) {
	cfg.T().Logf("[AfterStep] step finished: %s", name)
}
