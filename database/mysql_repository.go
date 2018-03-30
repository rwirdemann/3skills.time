package database

import (
	"github.com/rwirdemann/3skills.time/domain"
)

type MySQLRepository struct {
	projects map[int]domain.Project
	bookings map[domain.Project][]domain.Booking
}

func NewMySQLRepository() *MySQLRepository {
	r := MySQLRepository{projects: make(map[int]domain.Project), bookings: make(map[domain.Project][]domain.Booking)}
	r.AddProject(domain.Project{Name: "Picue"})
	r.AddProject(domain.Project{Name: "Energie"})

	b := domain.Booking{Description: "NRG-213", Hours: 2.0, ProjectId: 1}
	r.AddBooking(b)

	return &r
}

func (this *MySQLRepository) AllProjects(filter string) []domain.Project {
	result := []domain.Project{}
	for _, v := range this.projects {
		result = append(result, v)
	}
	return result
}

func (this *MySQLRepository) AllBookings(projectId int) []domain.Booking {
	return this.bookings[this.projects[projectId]]
}

func (this *MySQLRepository) AddProject(p domain.Project) {
	p.Id = this.nextProjectId()
	this.projects[p.Id] = p
}

func (this *MySQLRepository) AddBooking(b domain.Booking) {
	b.Id = this.nextBookingId()
	project := this.projects[b.ProjectId]
	this.bookings[project] = append(this.bookings[project], b)
}

func (this *MySQLRepository) nextProjectId() int {
	nextId := 1
	for k, _ := range this.projects {
		if k >= nextId {
			nextId = k + 1
		}
	}
	return nextId
}

func (this *MySQLRepository) nextBookingId() int {
	nextId := 1
	for _, bookings := range this.bookings {
		for _, b := range bookings {
			if b.Id >= nextId {
				nextId = b.Id + 1
			}
		}
	}
	return nextId
}
