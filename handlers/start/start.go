package start

import (
	"net/http"

	"github.com/chonla/umock/helpers"
)

func (h *StartHandler) Start() error {
	h.initializeHandler()

	h.log.Trace("Starting server at %s ...\n", h.conf.Server.String())
	if err := http.ListenAndServe(h.conf.Server.String(), nil); err != nil {

	}
	return nil
}

func (h *StartHandler) initializeHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, route := range h.conf.Routes {
			h.log.Debug("Matching %s ...", helpers.StringIfEmpty(route.Name, "Unnamed"))
			if route.Match(r, h.log) {
				pathWithQuery := r.URL.Path
				queryString := r.URL.Query().Encode()
				if queryString != "" {
					pathWithQuery = pathWithQuery + "?" + queryString
				}
				h.log.Trace("%s %s\n", r.Method, pathWithQuery)
				if route.Then != nil {
					h.log.Debug("Then found, apply Then ...")
					route.Then.Respond(w)
				}

				return
			}
		}

		w.WriteHeader(404)
	})
}
