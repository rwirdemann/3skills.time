package usecase

import (
	"log"

	"github.com/rwirdemann/go-tracker/design/domain"
	"github.com/rwirdemann/go-tracker/foundation"
)

type AddProject struct {
	consumer   foundation.Consumer
	repository Repository
}

func NewAddProject(consumer foundation.Consumer, repository Repository) *AddProject {
	return &AddProject{consumer: consumer, repository: repository}
}

func (this AddProject) Run(i interface{}) interface{} {
	project := this.consumer.Consume(i).(domain.Project)
	log.Printf("AddProject.Run: %s", project.Name)
	this.repository.Add(project)
	return project
}
