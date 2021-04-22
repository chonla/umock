package models

import (
	"net/http"

	"github.com/chonla/umock/logger"
)

type Route struct {
	Name   string `yml:"name,omitempty"`
	Method Method `yml:"method"`
	Path   Path   `yml:"path"`
	When   *When  `yml:"when,omitempty"`
	Then   *Then  `yml:"then,omitempty"`
}

func (o Route) Match(r *http.Request, log *logger.Logger) bool {
	if !o.Method.Test(r.Method) {
		// Method does not match
		log.Debug("  Matching Method ... FAILED")
		return false
	}
	log.Debug("  Matching Method ... PASSED")

	if !o.Path.Test(r.URL.Path) {
		// Path does not match
		log.Debug("  Matching Path ... FAILED")
		return false
	}
	log.Debug("  Matching Path ... PASSED")

	if o.When == nil {
		// No   matching condition is required, TRUE if path and method match
		log.Debug("  Matching When ... PASSED (No When is given)")
		return true
	}

	if !o.When.Test(r, log) {
		// Condition does not meet
		log.Debug("  Matching When ... FAILED")
		return false
	}

	return true
}
