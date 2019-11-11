package config

import (
	"strings"
)

type Replacer struct {
	*strings.Replacer
}

func NewReplacer(oldnew ...string) *Replacer {
	return &Replacer{
		Replacer: strings.NewReplacer(oldnew...),
	}
}
