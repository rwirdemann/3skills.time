package usecase

import (
	"log"

	"github.com/rwirdemann/gotracker-pg/foundation"
	"github.com/rwirdemann/gotracker/domain"
)

type AddProject struct {
	consumer   foundation.Consumer
	repository Repository
}

func NewAddProject(consumer foundation.Consumer, repository Repository) *AddProject {
	return &AddProject{consumer: consumer, repository: repository}
}

func (this AddProject) Run(i ...interface{}) interface{} {
	project := this.consumer.Consume(i[0]).(*domain.Project)
	log.Printf("AddProject.Run: %s", project.Name)
	this.repository.Add(*project)
	return project
}
