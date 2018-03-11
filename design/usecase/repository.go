package usecase

import "github.com/rwirdemann/go-tracker/design/domain"

type Repository interface {
	AllProjects(filter string) []domain.Project
	Add(p domain.Project)
}
