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

func (a *Assertion) NoError(err error) {
	a.cfg.Step("Check that request completed without transport error", func() {
		a.require.NoError(err)
	})
}

func (a *Assertion) Equal(actual, expected any, name string) {
	a.cfg.Step(fmt.Sprintf("Check that %s equals to %v", name, expected), func() {
		a.require.Equal(expected, actual, "Incorrect value: %s", name)
	})
}

func (a *Assertion) Nil(actual any, name string) {
	a.cfg.Step(fmt.Sprintf("Check that %s is nil", name), func() {
		a.require.Nil(actual, "Value must be nil: %s", name)
	})
}

func (a *Assertion) NotNil(actual any, name string) {
	a.cfg.Step(fmt.Sprintf("Check that %s is not nil", name), func() {
		a.require.NotNil(actual, "Value must not be nil: %s", name)
	})
}

func (a *Assertion) NotEmpty(actual any, name string) {
	a.cfg.Step(fmt.Sprintf("Check that %s is not empty", name), func() {
		a.require.NotEmpty(actual, "Value must not be empty: %s", name)
	})
}

func (a *Assertion) Error(actual error, name string) {
	a.cfg.Step(fmt.Sprintf("Check that %s contains an error", name), func() {
		a.require.Error(actual, "Expected an error: %s", name)
	})
}

func (a *Assertion) Fail(message string) {
	a.cfg.Step(fmt.Sprint(message), func() {
		a.require.Fail(message)
	})
}
