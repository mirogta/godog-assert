package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type GodogAssert struct {
	assert  *assert.Assertions
	require *require.Assertions
}

func NewGodogAssert(bt *BddTesting) *GodogAssert {
	return &GodogAssert{
		assert:  assert.New(bt),
		require: require.New(bt),
	}
}
