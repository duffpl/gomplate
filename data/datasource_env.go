package data

import (
	"strings"

	"github.com/duffpl/gomplate/v3/env"
)

func readEnv(source *Source, args ...string) (b []byte, err error) {
	n := source.URL.Path
	n = strings.TrimPrefix(n, "/")
	if n == "" {
		n = source.URL.Opaque
	}

	b = []byte(env.Getenv(n))
	return b, nil
}
