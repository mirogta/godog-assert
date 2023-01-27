package main

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type godogs struct {
	*GodogAssert
	available int
}

func NewGodogs(t *testing.T) *godogs {
	return &godogs{
		available:   1,
		GodogAssert: NewGodogAssert(t),
	}
}

func (g *godogs) thereAreGodogs(available int) {
	g.available = available
}

func (g *godogs) iEat(num int) {
	require.NotEmpty(g.t, g.available, "there are no godogs available")
	require.Less(g.t, num, g.available, fmt.Sprintf("you cannot eat %d godogs, there are %d available", num, g.available))

	g.available -= num
}

func (g *godogs) thereShouldBeRemaining(remaining int) {
	require.NotEmpty(g.t, g.available, "there are no godogs available")

	assert.Equal(g.t, g.available, remaining, fmt.Sprintf("expected %d godogs to be remaining, but there is %d", remaining, g.available))
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			InitializeScenario(t, sc)
		},
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

func InitializeScenario(t *testing.T, sc *godog.ScenarioContext) {
	g := NewGodogs(t)

	sc.Step(`^there are (\d+) godogs$`, g.thereAreGodogs)
	sc.Step(`^I eat (\d+)$`, g.iEat)
	sc.Step(`^there should be (\d+) remaining$`, g.thereShouldBeRemaining)
}
