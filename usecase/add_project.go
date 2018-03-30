package usecase

import (
	"log"

	"github.com/rwirdemann/gotracker-pg/foundation"
	"github.com/rwirdemann/3skills.time/domain"
)

type AddProject struct {
	consumer   foundation.Consumer
	repository Repository
}

func NewAddProject(consumer foundation.Consumer, repository Repository) *AddProject {
	return &AddProject{consumer: consumer, repository: repository}
}

func (u AddProject) Run(i ...interface{}) interface{} {
	project := u.consumer.Consume(i[0]).(*domain.Project)
	log.Printf("AddProject.Run: %s", project.Name)
	u.repository.AddProject(*project)
	return project
}
