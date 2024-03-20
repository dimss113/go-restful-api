package test

import (
	"dimasfadilah/go-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService(true)

	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)

	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)
	cleanup()
}
