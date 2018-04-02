package database

import (
	"github.com/rwirdemann/3skills.time/domain"
)

type MySQLRepository struct {
	projects map[int]domain.Project
	bookings map[domain.Project][]domain.Booking
}

func (m *MySQLRepository) AddBookings(b ...domain.Booking) {
	panic("implement me")
}

func (m *MySQLRepository) AddCustomer(c domain.Customer) int {
	panic("implement me")
}

func NewMySQLRepository() *MySQLRepository {
	r := MySQLRepository{projects: make(map[int]domain.Project), bookings: make(map[domain.Project][]domain.Booking)}
	r.AddProject(domain.Project{Name: "Picue"})
	r.AddProject(domain.Project{Name: "Energie"})

	b := domain.Booking{Description: "NRG-213", Hours: 2.0, ProjectId: 1}
	r.AddBooking(b)

	return &r
}

func (m *MySQLRepository) AllProjects(filter string) []domain.Project {
	var result []domain.Project
	for _, v := range m.projects {
		result = append(result, v)
	}
	return result
}

func (m *MySQLRepository) AllBookings(projectId int) []domain.Booking {
	return m.bookings[m.projects[projectId]]
}

func (m *MySQLRepository) AddProject(p domain.Project) int {
	p.Id = m.nextProjectId()
	m.projects[p.Id] = p
	return p.Id
}

func (m *MySQLRepository) AddBooking(b domain.Booking) {
	b.Id = m.nextBookingId()
	project := m.projects[b.ProjectId]
	m.bookings[project] = append(m.bookings[project], b)
}

func (m *MySQLRepository) nextProjectId() int {
	nextId := 1
	for k := range m.projects {
		if k >= nextId {
			nextId = k + 1
		}
	}
	return nextId
}

func (m *MySQLRepository) AllBookingsByMonthAndYear(customerId int, month int, year int) []domain.Booking {
	panic("implement me")
}

func (m *MySQLRepository) InvoiceByAndMonthAndYear(customerId int, month int, year int) (domain.Invoice, bool) {
	panic("implement me")
}

func (m *MySQLRepository) nextBookingId() int {
	nextId := 1
	for _, bookings := range m.bookings {
		for _, b := range bookings {
			if b.Id >= nextId {
				nextId = b.Id + 1
			}
		}
	}
	return nextId
}
