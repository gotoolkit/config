package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := `{ "key": "value" }`
	config, err := New(WithReader(strings.NewReader(s), JSON))
	assert.NoError(t, err)
	fmt.Println(config.AllKeys())
	v := config.GetString("key")
	fmt.Println(v)
	assert.NotEmpty(t, v)
}
