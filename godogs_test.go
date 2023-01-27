package main

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

type godogs struct {
	*GodogAssert
	available int
}

func NewGodogs(bt *BddTesting) *godogs {
	return &godogs{
		available:   0,
		GodogAssert: NewGodogAssert(bt),
	}
}

func (g *godogs) thereAreGodogs(available int) {
	g.available = available
}

func (g *godogs) iEat(num int) {
	g.require.NotEmpty(g.available, "there are no godogs available")
	g.require.Less(num, g.available, fmt.Sprintf("you cannot eat %d godogs, there are %d available", num, g.available))

	g.available -= num
}

func (g *godogs) thereShouldBeRemaining(remaining int) {
	g.require.NotEmpty(g.available, "there are no godogs available")

	g.assert.Equal(g.available, remaining, fmt.Sprintf("expected %d godogs to be remaining, but there is %d", remaining, g.available))
}

func TestFeatures(t *testing.T) {
	bt := NewBddTesting(t)
	suite := godog.TestSuite{
		ScenarioInitializer: TestingScenarioInitialiser(bt, InitializeScenario),
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(bt *BddTesting, sc *godog.ScenarioContext) {
	g := NewGodogs(bt)

	sc.Step(`^there are (\d+) godogs$`, g.thereAreGodogs)
	sc.Step(`^I eat (\d+)$`, g.iEat)
	sc.Step(`^there should be (\d+) remaining$`, g.thereShouldBeRemaining)
}
