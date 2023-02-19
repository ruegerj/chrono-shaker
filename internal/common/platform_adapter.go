package common

import (
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type PlatformAdapter interface {
	CreateListingsUrl() string
	Parse(g *geziyor.Geziyor, r *client.Response)
}
