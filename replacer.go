package config

import (
	"strings"
)

func NewReplacer(oldnew ...string) *strings.Replacer {
	return strings.NewReplacer(oldnew...)
}
