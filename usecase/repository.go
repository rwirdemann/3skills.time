package usecase

import "github.com/rwirdemann/3skills.time/domain"

type Repository interface {
	AddCustomer(c domain.Customer) int

	AllProjects(filter string) []domain.Project
	AddProject(p domain.Project) int
	AddBooking(b domain.Booking)
	AllBookings(projectId int) []domain.Booking

	AllBookingsByMonthAndYear(customerId int, month int, year int) []domain.Booking
	InvoiceByAndMonthAndYear(customerId int, month int, year int) (domain.Invoice, bool)
}
