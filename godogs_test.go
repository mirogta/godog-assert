package main

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type godogs struct {
	bt        *BddTesting
	available int
}

func NewGodogs(bt *BddTesting) *godogs {
	return &godogs{
		available: 0,
		bt:        bt,
	}
}

func (g *godogs) thereAreGodogs(available int) {
	g.available = available
}

func (g *godogs) iEat(num int) {
	require.NotEmpty(g.bt, g.available, "there are no godogs available")
	require.Less(g.bt, num, g.available, fmt.Sprintf("you cannot eat %d godogs, there are %d available", num, g.available))

	g.available -= num
}

func (g *godogs) thereShouldBeRemaining(remaining int) {
	require.NotEmpty(g.bt, g.available, "there are no godogs available")

	assert.Equal(g.bt, g.available, remaining, fmt.Sprintf("expected %d godogs to be remaining, but there is %d", remaining, g.available))
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: TestingScenarioInitialiser(t, InitializeScenario),
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

func InitializeScenario(t *testing.T, sc *godog.ScenarioContext) *BddTesting {
	bt := NewBddTesting(t)
	g := NewGodogs(bt)

	sc.Step(`^there are (\d+) godogs$`, g.thereAreGodogs)
	sc.Step(`^I eat (\d+)$`, g.iEat)
	sc.Step(`^there should be (\d+) remaining$`, g.thereShouldBeRemaining)

	return bt
}
