package database

import (
	"github.com/rwirdemann/go-tracker/design/domain"
)

type MySQLRepository struct {
}

func NewMySQLRepository() *MySQLRepository {
	return &MySQLRepository{}
}

func (this MySQLRepository) AllProjects(filter string) []domain.Project {
	picue := domain.Project{Name: "Picue"}
	energy := domain.Project{Name: "Energie"}
	return []domain.Project{picue, energy}
}

func (this MySQLRepository) Add(p domain.Project) {

}
