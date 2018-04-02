package usecase

import (
	"github.com/rwirdemann/3skills.time/domain"
	"time"
)

type FakeRepository struct {
	customers         map[int]domain.Customer
	projects          map[int]domain.Project
	bookingsByProject map[domain.Project][]domain.Booking

	invoices map[int]map[int]map[int]domain.Invoice // invoices[customerid][year][month]
}

func (f *FakeRepository) AddBookings(bookings ...domain.Booking) {
	for _, b := range bookings {
		f.AddBooking(b)
	}
}

func NewFakeRepository() *FakeRepository {
	r := FakeRepository{projects: make(map[int]domain.Project),
		bookingsByProject: make(map[domain.Project][]domain.Booking),
		customers: make(map[int]domain.Customer),
		invoices: make(map[int]map[int]map[int]domain.Invoice)}

	return &r
}

func (f *FakeRepository) AddCustomer(c domain.Customer) int {
	c.Id = f.nextCustomerId()
	f.customers[c.Id] = c
	return c.Id
}

func (f *FakeRepository) AllProjects(filter string) []domain.Project {
	var result []domain.Project
	for _, v := range f.projects {
		result = append(result, v)
	}
	return result
}

func (f *FakeRepository) AllBookings(projectId int) []domain.Booking {
	return f.bookingsByProject[f.projects[projectId]]
}

func (f *FakeRepository) AddProject(p domain.Project) int {
	p.Id = f.nextProjectId()
	f.projects[p.Id] = p
	return p.Id
}

func (f *FakeRepository) AddBooking(b domain.Booking) {
	b.Id = f.nextBookingId()
	project := f.projects[b.ProjectId]
	f.bookingsByProject[project] = append(f.bookingsByProject[project], b)
}

func (f *FakeRepository) AllBookingsByMonthAndYear(customerId int, year int, month int) []domain.Booking {
	var result []domain.Booking
	for project, bookings := range f.bookingsByProject {
		if project.CustomerId == customerId {
			for _, b := range bookings {
				if b.Date.Year() == year && b.Date.Month() == time.Month(month) {
					result = append(result, b)
				}
			}
		}
	}
	return result
}

func (f *FakeRepository) InvoiceByAndMonthAndYear(customerId int, year int, month int) (domain.Invoice, bool) {
	if _, ok := f.invoices[customerId]; !ok {
		m := make(map[int]map[int]domain.Invoice)
		f.invoices[customerId] = m
	}

	if _, ok := f.invoices[customerId][year]; !ok {
		m := make(map[int]domain.Invoice)
		f.invoices[customerId][year] = m
	}

	if i, ok := f.invoices[customerId][year][month]; ok {
		return i, true
	}

	return domain.Invoice{}, false
}

func (f *FakeRepository) nextProjectId() int {
	nextId := 1
	for k := range f.projects {
		if k >= nextId {
			nextId = k + 1
		}
	}
	return nextId
}

func (f *FakeRepository) nextCustomerId() int {
	nextId := 1
	for k := range f.customers {
		if k >= nextId {
			nextId = k + 1
		}
	}
	return nextId
}

func (f *FakeRepository) nextBookingId() int {
	nextId := 1
	for _, bookings := range f.bookingsByProject {
		for _, b := range bookings {
			if b.Id >= nextId {
				nextId = b.Id + 1
			}
		}
	}
	return nextId
}
func (f *FakeRepository) contains(name string) bool {
	return true
}
