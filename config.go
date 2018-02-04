package cmf

import (
	"encoding/json"
	"golang.org/x/net/idna"
	"golang.org/x/text/language"
)

type config struct {
	database string
	// The first language in this slice is the default. Listed in order from
	// most supported to least.
	languages []language.Tag
	// The matcher initialized with languages.
	matcher language.Matcher
	domains []domain
	// The domain used for API requests. If the domain already exists in
	// domains, the API paths will be prefixed with /api/.
	api string
}

func readConfig(dir string) (*config, error) {
	conf = new(config)
	conf.matcher = language.NewMatcher(conf.languages)
	return conf, nil
}

func (c *config) writeConfig(path string) error {

}

type domain struct {
	// All domain names are stored in their ASCII-compatible format.
	name string
	// Domains can optionally specify a locale. ccTLDs should almost always
	// specify a region.
	// e.g. example.ca might specify und-CA, whereas exemple.ca or fr.example.ca
	// might specify fr-CA.
	// Further, a script should be specified when a domain should only serve
	// content in said script, e.g. a domain written in zh-Hans should probably
	// not serve content in zh-Hant.
	language.Tag
}
