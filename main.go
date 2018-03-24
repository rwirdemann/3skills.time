package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"github.com/rwirdemann/gotracker/domain"
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
	presenter := rest.NewJSONPresenter()
	repository := database.NewMySQLRepository()
	getProjects := usecase.NewGetProjects(consumer, presenter, repository)

	projectConsumer := rest.NewJSONConsumer(&domain.Project{})
	addProject := usecase.NewAddProject(projectConsumer, repository)

	projectIdConsumer := rest.NewURLConsumer("projectId", "int")
	getBookings := usecase.NewGetBookings(projectIdConsumer, presenter, repository)

	bookingConsumer := rest.NewJSONConsumer(&domain.Booking{})
	addBooking := usecase.NewAddBooking(projectIdConsumer, bookingConsumer, repository)

	r := mux.NewRouter()
	getProjectsHandler := rest.MakeGetProjectsHandler(getProjects)
	addProjectHandler := rest.MakeAddProjectHandler(addProject)
	getBookingsHandler := rest.MakeGetBookingsHandler(getBookings)
	addBookingHandler := rest.MakeAddBookingHandler(addBooking)

	if *unsecure {
		r.HandleFunc("/projects", getProjectsHandler).Methods("GET")
		r.HandleFunc("/projects", addProjectHandler).Methods("POST")
		r.HandleFunc("/projects/{projectId}/bookings", getBookingsHandler).Methods("GET")
		r.HandleFunc("/projects/{projectId}/bookings", addBookingHandler).Methods("POST")
	} else {
		r.HandleFunc("/projects", middleware.JWT(getProjectsHandler)).Methods("GET")
		r.HandleFunc("/projects", middleware.JWT(addProjectHandler)).Methods("POST")
	}

	fmt.Println("GET  http://localhost:8080/projects")
	fmt.Println("POST http://localhost:8080/projects")
	fmt.Println("GET  http://localhost:8080/projects/1/bookings")
	fmt.Println("POST http://localhost:8080/projects/1/bookings")

	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
}
