package models

import (
	"net/http"
)

type Route struct {
	Method Method `yml:"method"`
	Path   Path   `yml:"path"`
	When   *When  `yml:"when,omitempty"`
	Then   *Then  `yml:"then,omitempty"`
}

func (o Route) Match(r *http.Request) bool {
	if !o.Method.Test(r.Method) {
		// Method does not match
		return false
	}

	if !o.Path.Test(r.URL.Path) {
		// Path does not match
		return false
	}

	if o.When == nil {
		// No matching condition is required, TRUE if path and method match
		return true
	}

	if !o.When.Test(r) {
		// Condition does not meet
		return false
	}

	return true
}
