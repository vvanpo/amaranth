package cmf

import (
	"regexp"
)

type Router interface {
	Route(path string) (Resource, error)
	Register(identifier string, param regexp.Regexp, r Resource) error
	Deregister(identifier string)
}
