package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

type BddTesting struct {
	t   *testing.T
	err error
}

func NewBddTesting(t *testing.T) *BddTesting {
	return &BddTesting{t: t}
}

func (bt *BddTesting) Testing() *testing.T {
	return bt.t
}

func (bt *BddTesting) Errorf(format string, args ...interface{}) {
	bt.err = fmt.Errorf(format, args...)

	// do not log additional error message - it's already too chatty
	//t.test.Errorf(format, args...)

	bt.t.Fail()
}

func (bt *BddTesting) FailNow() {
	// actually, don't FailNow bacause that will panic
	// and prevent from getting the overall suite stats
	bt.t.Fail()
}

func (bt *BddTesting) EmitErrors(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
	if bt.err == nil {
		return ctx, nil
	}

	// TODO: there is still some pending issue here…
	// even though godog now recognizes require/assert failures
	// it duplicates the error output
	// once (correctly) in a red colour after the failed scenario
	// but then another one (incorrectly) in a grey colour, triggered at suite.go:451

	return ctx, bt.err
}

// Wrapper around ScenarioInitialiser which allows us to pass the *testing.T object to it along with the ScenarioContext
func TestingScenarioInitialiser(t *testing.T, init func(*testing.T, *godog.ScenarioContext) *BddTesting) func(sc *godog.ScenarioContext) {
	return func(sc *godog.ScenarioContext) {
		bt := init(t, sc)
		// this produces the nice color-coded output
		sc.StepContext().After(bt.EmitErrors)
	}
}
