package database

import (
	"github.com/rwirdemann/go-tracker/design/domain"
)

func NewMySQLProjectRepository() *MySQLProjectRepository {
	return &MySQLProjectRepository{}
}

type MySQLProjectRepository struct {
}

func (this MySQLProjectRepository) All() []domain.Project {
	picue := domain.Project{Name: "Picue"}
	energy := domain.Project{Name: "Energie"}
	return []domain.Project{picue, energy}
}
