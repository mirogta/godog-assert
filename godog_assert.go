package main

import (
	"testing"
)

type GodogAssert struct {
	t *testing.T
}

func NewGodogAssert(t *testing.T) *GodogAssert {
	return &GodogAssert{
		t: t,
	}
}
