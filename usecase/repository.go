package usecase

import "github.com/rwirdemann/gotracker/domain"

type Repository interface {
	AllProjects(filter string) []domain.Project
	Add(p domain.Project)
	AddBooking(b domain.Booking)
	AllBookings(projectId int) []domain.Booking
}
