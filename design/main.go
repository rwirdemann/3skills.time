package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rwirdemann/go-tracker/design/database"
	"github.com/rwirdemann/go-tracker/design/rest"
	"github.com/rwirdemann/go-tracker/design/usecase"
)

func main() {
	presenter := rest.NewJSONPresenter()
	repository := database.NewMySQLProjectRepository()

	r := mux.NewRouter()
	r.HandleFunc("/projects", rest.MakeGetProjectsHandler(usecase.NewGetProjects(presenter, repository)))
	http.ListenAndServe(":8080", r)
}
