package driver

import (
	"github.com/cwd-nial/cors-vs-json-p/cors"
	"github.com/cwd-nial/cors-vs-json-p/jsonp"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter constructs new mux.Router.
func NewRouter(h Handler) *mux.Router {
	r := mux.NewRouter()

	setupRoutes(r, h)
	setupJsonP(r)
	setupCors(r)

	return r
}

// setupRoutes setups REST-API routes.
func setupRoutes(r *mux.Router, h Handler) {
	r.Methods(http.MethodGet).Path("/info").HandlerFunc(h.GetInfo())
}

func setupJsonP(r *mux.Router) {
	h := http.FileServer(http.FS(jsonp.GetStaticFiles()))
	r.PathPrefix("/jsonp/").Handler(http.StripPrefix("/jsonp/", h))
}

func setupCors(r *mux.Router) {
	h := http.FileServer(http.FS(cors.GetStaticFiles()))
	r.PathPrefix("/cors/").Handler(http.StripPrefix("/cors/", h))
}
