package usecase

import (
	"github.com/rwirdemann/gotracker/domain"
)

type FakeRepository struct {
	projects map[int]domain.Project
	bookings map[domain.Project][]domain.Booking
}

func NewFakeRepository() *FakeRepository {
	r := FakeRepository{projects: make(map[int]domain.Project), bookings: make(map[domain.Project][]domain.Booking)}
	r.AddProject(domain.Project{Name: "Picue"})
	r.AddProject(domain.Project{Name: "Energie"})

	b := domain.Booking{Description: "NRG-213", Hours: 2.0, ProjectId: 1}
	r.AddBooking(b)

	return &r
}

func (this *FakeRepository) contains(name string) bool {
	return true
}

func (this *FakeRepository) AllProjects(filter string) []domain.Project {
	result := []domain.Project{}
	for _, v := range this.projects {
		result = append(result, v)
	}
	return result
}

func (this *FakeRepository) AllBookings(projectId int) []domain.Booking {
	return this.bookings[this.projects[projectId]]
}

func (this *FakeRepository) AddProject(p domain.Project) {
	p.Id = this.nextProjectId()
	this.projects[p.Id] = p
}

func (this *FakeRepository) AddBooking(b domain.Booking) {
	b.Id = this.nextBookingId()
	project := this.projects[b.ProjectId]
	this.bookings[project] = append(this.bookings[project], b)
}

func (this *FakeRepository) nextProjectId() int {
	nextId := 1
	for k, _ := range this.projects {
		if k >= nextId {
			nextId = k + 1
		}
	}
	return nextId
}

func (this *FakeRepository) nextBookingId() int {
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
