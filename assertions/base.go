package assertions

import (
	"fmt"

	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/require"
)

type Assertion struct {
	cfg     *axiom.Config
	require *require.Assertions
}

func New(cfg *axiom.Config) *Assertion {
	return &Assertion{
		cfg:     cfg,
		require: require.New(cfg.T()),
	}
}

func (a *Assertion) NoError(actual error) {
	message := "Check that request completed without transport error"
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.require.NoError(actual)
	})
}

func (a *Assertion) Equal(actual, expected any, name string) {
	message := fmt.Sprintf("Check that %s equals to %v", name, expected)
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.require.Equal(expected, actual, "Incorrect value: %s", name)
	})
}

func (a *Assertion) Nil(actual any, name string) {
	message := fmt.Sprintf("Check that %s is nil", name)
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.require.Nil(actual, "Value must be nil: %s", name)
	})
}

func (a *Assertion) NotNil(actual any, name string) {
	message := fmt.Sprintf("Check that %s is not nil", name)
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.require.NotNil(actual, "Value must not be nil: %s", name)
	})
}

func (a *Assertion) NotEmpty(actual any, name string) {
	message := fmt.Sprintf("Check that %s is not empty", name)
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.require.NotEmpty(actual, "Value must not be empty: %s", name)
	})
}

func (a *Assertion) Error(actual error, name string) {
	message := fmt.Sprintf("Check that %s contains an error", name)
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.require.Error(actual, "Expected an error: %s", name)
	})
}

func (a *Assertion) Fail(message string) {
	a.cfg.Step(message, func() {
		a.cfg.Log(axiom.NewInfoLog(message))
		a.require.Fail(message)
	})
}
