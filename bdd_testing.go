package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

type BddTesting struct {
	test *testing.T
	err  error
}

func NewBddTesting(t *testing.T) *BddTesting {
	return &BddTesting{
		test: t,
	}
}

func (t *BddTesting) Testing() *testing.T {
	return t.test
}

func (t *BddTesting) Errorf(format string, args ...interface{}) {
	t.err = fmt.Errorf(format, args...)

	// do not log additional error message - it's already too chatty
	//t.test.Errorf(format, args...)

	t.test.Fail()
}

func (t *BddTesting) FailNow() {
	t.test.FailNow()
}

func (t *BddTesting) EmitErrors(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
	if t.err == nil {
		return ctx, nil
	}

	// TODO: there is still some pending issue here…
	// even though godog now recognizes require/assert failures
	// it duplicates the error output
	// once (correctly) in a red colour after the failed scenario
	// but then another one (incorrectly) in a grey colour, triggered at suite.go:451

	return ctx, t.err
}

// Wrapper around ScenarioInitialiser which allows us to pass the BddTesting object to it along with the ScenarioContext
func TestingScenarioInitialiser(bt *BddTesting, init func(*BddTesting, *godog.ScenarioContext)) func(sc *godog.ScenarioContext) {
	return func(sc *godog.ScenarioContext) {
		sc.StepContext().After(bt.EmitErrors)
		init(bt, sc)
	}
}
