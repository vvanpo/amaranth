package cmf

import (
	"golang.org/x/text/language"
	"net/http"
	"time"
)

// The exchange object manages state for a single HTTP request/response pair.
type exchange struct {
	*Server
	*http.Request
	http.ResponseWriter
	*user // nil if not signed in.
	// This tag may be different than the associated supported language, see
	// https://godoc.org/golang.org/x/text/language#hdr-Using_match_results
	locale language.Tag
}
