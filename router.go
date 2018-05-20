package cmf

import (
	"regexp"
)

type Router interface {
	Route(path string) (Resource, error)
}

type StaticRouter struct {
	paths map[string]Resource
}

func (s *StaticRouter) Route(path string) (Resource, error) {

}

func (s *StaticRouter) Register(identifier string, r Resource) error {

}

func (s *StaticRouter) Deregister(identifier string) {

}

type paramRouter struct {
	name  string
	param *regexp.Regexp
	r     Resource
}

func (p *paramRouter) Route(path string) (Resource, error) {

}

func ParamRouter(name string, regexp string, r Resource) (Router, error) {
	param, err := regexp.Compile(regexp)
	if err != nil {
		return nil, err
	}
	return &paramRouter{name, param.Copy(), r}, nil
}
