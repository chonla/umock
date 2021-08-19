package models

import (
	"net/http"
	"strings"
)

type Then struct {
	Status  int      `yml:"status,omitempty"`
	Headers []string `yml:"headers,omitempty"`
	Body    string   `yml:"body,omitempty"`
}

func (t Then) Respond(w http.ResponseWriter) {
	if len(t.Headers) > 0 {
		for _, h := range t.Headers {
			hPair := strings.SplitN(h, "=", 2)
			w.Header().Add(hPair[0], hPair[1])
		}
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if t.Status > 0 {
		w.WriteHeader(t.Status)
	}
	if len(t.Body) > 0 {
		w.Write([]byte(t.Body))
	}
}
