package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"github.com/rwirdemann/gotracker/middleware"

	"github.com/gorilla/mux"
	"github.com/rwirdemann/gotracker/database"
	"github.com/rwirdemann/gotracker/rest"
	"github.com/rwirdemann/gotracker/usecase"
)

func main() {
	unsecure := flag.Bool("unsecure", false, "run in unsecure mode")
	flag.Parse()

	consumer := rest.NewQueryConsumer()
	jsonConsumter := rest.NewJSONConsumer()
	presenter := rest.NewJSONPresenter()
	repository := database.NewMySQLRepository()
	getProjects := usecase.NewGetProjects(consumer, presenter, repository)
	addProject := usecase.NewAddProject(jsonConsumter, repository)

	r := mux.NewRouter()
	getProjectsHandler := rest.MakeGetProjectsHandler(getProjects)
	addProjectHandler := rest.MakeAddProjectHandler(addProject)
	if *unsecure {
		r.HandleFunc("/projects", getProjectsHandler).Methods("GET")
		r.HandleFunc("/projects", addProjectHandler).Methods("POST")
	} else {
		r.HandleFunc("/projects", middleware.JWT(getProjectsHandler)).Methods("GET")
		r.HandleFunc("/projects", middleware.JWT(addProjectHandler)).Methods("POST")
	}

	fmt.Println("GET  http://localhost:8080/projects")
	fmt.Println("POST http://localhost:8080/projects")

	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
}
