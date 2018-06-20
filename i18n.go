package cmf

import (
	"context"
	"net/http"
)

/*
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Attempt to match a supported language.
	locale := matchLocale(req, s.conf) // Parse the request URL and determine if it is valid, return a 400-level
	// error otherwise.
	// If more than one language is configured, yet the requested URL does not
	// specify one, attempt to match a supported language using the
	// Accept-Language header.
	if locale == language.Und && len(s.conf.Languages) > 1 {
		locale, i = language.MatchStrings(s.conf.Matcher, req.Header.Get("Accept-Language"))
	}
	// If the matched language isn't the default, redirect to the matched
	// language provided the resource is valid.
	matchedLanguage = s.conf.Languages[i]

	// If the matched language above is different from the requested default
	// URL, redirect to the correct language prefix using path translation if it
	// is enabled.
	w.Header().Set("Location")
	w.WriteHeader(http.StatusTemporaryRedirect)
	//
	e := exchange{s, req, w}
}*/
