package models

import (
	"net/http"

	"github.com/chonla/umock/logger"
)

type When struct {
	ContentType ContentType `yaml:"content_type"`
	Param       Param       `yaml:"param"`
	Query       Query       `yaml:"query"`
	Form        Form        `yaml:"form"`
	JsonBody    JsonBody    `yaml:"json_body"`
}

func (w When) Test(r *http.Request, params map[string]string, log *logger.Logger) bool {
	if !w.Query.Test(r.URL.Query(), log) {
		log.Debug("  Matching Querystring ... FAILED")
		return false
	}
	log.Debug("  Matching Querystring ... PASSED")

	if !w.Param.Test(params, log) {
		log.Debug("  Matching Path Params ... FAILED")
		return false
	}
	log.Debug("  Matching Path Params ... PASSED")

	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		log.Debug("  Detected Content-Type: %s", r.Header.Get("Content-Type"))

		if !w.ContentType.Test(r.Header.Get("Content-Type")) {
			log.Debug("  Matching Content-Type ... FAILED")
			return false
		}
		log.Debug("  Matching Content-Type ... PASSED")

		if r.Header.Get("Content-Type") == "application/json" {
			if !w.JsonBody.Test(r, log) {
				log.Debug("  Matching JSON Body ... FAILED")
				return false
			}
			log.Debug("  Matching JSON Body ... PASSED")
		} else {
			err := r.ParseForm()
			if err != nil {
				// Form cannot be parsed
				log.Error("Matching Form Body ... %v", err)
				return false
			}
			if !w.Form.Test(r.PostForm, log) {
				log.Debug("  Matching Form Body ... FAILED")
				return false
			}
			log.Debug("  Matching Form Body ... PASSED")
		}
	}
	return true
}
