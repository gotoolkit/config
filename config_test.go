package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	err := Setup()
	host := GetString("database.host")
	log.Println(err)
	assert.NoError(t, err)
	assert.NotEmpty(t, host)
}
