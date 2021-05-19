package caster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCaster(t *testing.T) {
	res := NewCaster()
	assert.NotNil(t, res)
}
