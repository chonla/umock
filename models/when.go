package models

import "net/http"

type When struct {
	Query Query `yaml:"query"`
	Form  Form  `yaml:"form"`
}

func (w When) Test(r *http.Request) bool {
	if !w.Query.Test(r.URL.Query()) {
		return false
	}
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		e := r.ParseForm()
		if e != nil {
			// Form cannot be parsed
			return false
		}
		if !w.Form.Test(r.PostForm) {
			return false
		}
	}
	return true
}
