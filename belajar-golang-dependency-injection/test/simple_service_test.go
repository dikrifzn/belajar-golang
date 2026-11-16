package test

import (
	"belajar-golang-dependency-injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeSimpleService(true)

	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeSimpleService(false)

	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
