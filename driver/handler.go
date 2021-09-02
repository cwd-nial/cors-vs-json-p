package driver

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Handler interface {
	GetInfo() HandlerFunc
}

type handler struct {
}

func NewHandler() Handler {
	return handler{}
}

func (h handler) GetInfo() HandlerFunc {
	type Version struct {
		AppName  string `json:"app"`
		Revision string `json:"rev"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Set-Cookie", "xyz=test")

		err := json.NewEncoder(w).Encode(Version{
			AppName:  "cors-vs-json-p",
			Revision: "dev-0.0.1",
		})
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
