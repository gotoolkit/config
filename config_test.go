package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	err := Setup()
	host := GetString("database.host")

	assert.NoError(t, err)
	assert.NotEmpty(t, host)
}
