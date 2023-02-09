package main

import (
	"testing"

	"github.com/cucumber/godog"
)

type ScenarioInitialiser func(t *testing.T, sc *godog.ScenarioContext) *BddTesting

func NewDefaultGodogSuite(t *testing.T, init ScenarioInitialiser) *godog.TestSuite {
	suite := &godog.TestSuite{
		ScenarioInitializer: TestingScenarioInitialiser(t, init),
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features"},
		},
	}
	return suite
}

// Wrapper around ScenarioInitialiser which allows us to pass the *testing.T object to it along with the ScenarioContext
func TestingScenarioInitialiser(t *testing.T, init ScenarioInitialiser) func(sc *godog.ScenarioContext) {
	return func(sc *godog.ScenarioContext) {
		bt := init(t, sc)
		// this produces the nice color-coded output
		sc.StepContext().After(bt.EmitErrors)
	}
}
