package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rwirdemann/go-tracker/design/database"
	"github.com/rwirdemann/go-tracker/design/rest"
	"github.com/rwirdemann/go-tracker/design/usecase"
)

func main() {
	consumer := rest.NewQueryConsumer()
	presenter := rest.NewJSONPresenter()
	repository := database.NewMySQLRepository()
	getProjects := usecase.NewGetProjects(consumer, presenter, repository)

	r := mux.NewRouter()
	r.HandleFunc("/projects", rest.MakeGetProjectsHandler(getProjects)).Methods("GET")
	r.HandleFunc("/projects", rest.MakeAddProjectHandler(usecase.NewAddProject(consumer, repository))).Methods("POST")

	fmt.Println("GET  http://localhost:8080/projects")
	fmt.Println("POST http://localhost:8080/projects")

	http.ListenAndServe(":8080", r)
}
