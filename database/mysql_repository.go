package database

import (
	"fmt"

	"github.com/rwirdemann/gotracker/domain"
)

type MySQLRepository struct {
	projects map[int]domain.Project
}

func NewMySQLRepository() *MySQLRepository {
	r := MySQLRepository{projects: make(map[int]domain.Project)}
	r.Add(domain.Project{Name: "Picue"})
	r.Add(domain.Project{Name: "Energie"})
	return &r
}

func (this *MySQLRepository) AllProjects(filter string) []domain.Project {
	result := []domain.Project{}
	for _, v := range this.projects {
		result = append(result, v)
	}
	return result
}

func (this *MySQLRepository) Add(p domain.Project) {
	p.Id = this.nextId()
	this.projects[p.Id] = p
}

func (this *MySQLRepository) nextId() int {
	nextId := 1
	for k, _ := range this.projects {
		if k >= nextId {
			nextId = k + 1
		}
	}
	fmt.Printf("NextId %d", nextId)
	return nextId
}
