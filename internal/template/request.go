package template

import (
	"github.com/drewjya/forge/internal/util"
)

func Request(name string) string {
	N := util.Title(name)
	module := goModule()

	return `package request

import (
	"` + module + `/gen/models"
)

type ` + N + `Request struct {
	// TODO: add fields
}

func (req ` + N + `Request) ToSetter() models.` + N + `Setter {
	return models.` + N + `Setter{
		// TODO: map fields
	}
}
`
}
