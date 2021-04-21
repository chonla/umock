package start

import (
	"fmt"
	"net/http"
)

func (h *StartHandler) Start() error {
	h.initializeHandler()

	fmt.Printf("Starting server at %s ...\n", h.conf.Server.String())
	if err := http.ListenAndServe(h.conf.Server.String(), nil); err != nil {

	}
	return nil
}

func (h *StartHandler) initializeHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, route := range h.conf.Routes {
			if route.Match(r) {
				pathWithQuery := r.URL.Path
				queryString := r.URL.Query().Encode()
				if queryString != "" {
					pathWithQuery = pathWithQuery + "?" + queryString
				}
				fmt.Printf("%s %s\n", r.Method, pathWithQuery)

				if route.Then != nil {
					route.Then.Respond(w)
				}

				return
			}
		}

		w.WriteHeader(404)
	})
}
