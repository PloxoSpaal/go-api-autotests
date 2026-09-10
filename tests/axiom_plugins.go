package tests

import (
	"fmt"
	"time"

	"github.com/Nikita-Filonov/axiom"
)

func featureNamePlugin() axiom.Plugin {
	return func(cfg *axiom.Config) {
		cfg.Runtime.EmitTestWrap(func(next axiom.TestAction) axiom.TestAction {
			return func(cfg *axiom.Config) {

				defer func() {
					feature := cfg.Meta.Feature
					if feature == "" {
						feature = "Case"
					}
					cfg.Case.Name = fmt.Sprintf("[%s] %s", feature, cfg.Case.Name)
				}()

				next(cfg)
			}
		})
	}
}

func durationPlugin() axiom.Plugin {
	return func(cfg *axiom.Config) {
		cfg.Runtime.EmitTestWrap(func(next axiom.TestAction) axiom.TestAction {
			return func(cfg *axiom.Config) {
				startedAt := time.Now()

				defer func() {
					cfg.T().Logf(
						"[Plugin] test %q finished in %s",
						cfg.Case.Name,
						time.Since(startedAt),
					)
				}()

				next(cfg)
			}
		})

		cfg.Runtime.EmitStepWrap(func(name string, next axiom.StepAction) axiom.StepAction {
			return func() {
				startedAt := time.Now()

				defer func() {
					cfg.T().Logf(
						"[Plugin] step %q finished in %s",
						name,
						time.Since(startedAt),
					)
				}()

				next()
			}
		})
	}
}

func stepCounterPlugin() axiom.Plugin {
	return func(cfg *axiom.Config) {
		var counter int
		cfg.Runtime.EmitTestWrap(func(next axiom.TestAction) axiom.TestAction {
			return func(cfg *axiom.Config) {

				defer func() {
					cfg.T().Logf(
						"[Plugin] test \"%s\" executed %d steps",
						cfg.Case.Name,
						counter,
					)
				}()

				next(cfg)
			}
		})

		cfg.Runtime.EmitStepWrap(func(name string, next axiom.StepAction) axiom.StepAction {
			return func() {
				counter++
				defer func() {
					cfg.T().Logf(
						"[Plugin] step #%d started: %s",
						counter,
						name,
					)
				}()

				next()
			}
		})
	}
}
