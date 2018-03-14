package main

import (
	"fmt"
	"net/http"

	"github.com/rwirdemann/gotracker/middleware"

	"github.com/gorilla/mux"
	"github.com/rwirdemann/gotracker/database"
	"github.com/rwirdemann/gotracker/rest"
	"github.com/rwirdemann/gotracker/usecase"
)

func main() {
	consumer := rest.NewQueryConsumer()
	jsonConsumter := rest.NewJSONConsumer()
	presenter := rest.NewJSONPresenter()
	repository := database.NewMySQLRepository()
	getProjects := usecase.NewGetProjects(consumer, presenter, repository)
	addProject := usecase.NewAddProject(jsonConsumter, repository)

	r := mux.NewRouter()
	r.HandleFunc("/projects",
		middleware.JWT(
			rest.MakeGetProjectsHandler(getProjects))).Methods("GET")

	r.HandleFunc("/projects",
		middleware.JWT(
			rest.MakeAddProjectHandler(addProject))).Methods("POST")

	fmt.Println("GET  http://localhost:8080/projects")
	fmt.Println("POST http://localhost:8080/projects")

	http.ListenAndServe(":8080", r)
}
