package rest

import (
	"io/ioutil"
	"net/http"

	"github.com/rwirdemann/go-tracker/foundation"
)

func MakeGetProjectsHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(usecase.Run(r).([]byte))
	}
}

func MakeAddProjectHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		usecase.Run(body)
	}
}
