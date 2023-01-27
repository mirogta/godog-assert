package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type GodogAssert struct {
	assert  *assert.Assertions
	require *require.Assertions
}

type godogs struct {
	GodogAssert
	available int
}

func NewGodogs(t *testing.T) *godogs {
	return &godogs{
		available: 1,
		GodogAssert: GodogAssert{
			assert:  assert.New(t),
			require: require.New(t),
		},
	}
}

func (g *godogs) reset() {
	g.available = 0
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

	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		g.reset()
		return ctx, nil
	})

	sc.Step(`^there are (\d+) godogs$`, g.thereAreGodogs)
	sc.Step(`^I eat (\d+)$`, g.iEat)
	sc.Step(`^there should be (\d+) remaining$`, g.thereShouldBeRemaining)
}
