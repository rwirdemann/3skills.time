package rest

import (
	"io/ioutil"
	"net/http"

	"github.com/rwirdemann/go-tracker/design/usecase"
)

func MakeGetProjectsHandler(usecase *usecase.GetProjects) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(usecase.Run(nil).([]byte))
	}
}

func MakeAddProjectHandler(usecase *usecase.AddProject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		usecase.Run(body)
	}
}
