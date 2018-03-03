package rest

import (
	"net/http"

	"github.com/rwirdemann/go-tracker/design/usecase"
)

func MakeGetProjectsHandler(usecase *usecase.GetProjects) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(usecase.Run().([]byte))
	}
}
