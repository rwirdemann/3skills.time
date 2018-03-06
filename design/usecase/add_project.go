package usecase

import (
	"log"

	"github.com/rwirdemann/go-tracker/design/domain"
)

type AddProject struct {
	consumer   Consumer
	repository domain.ProjectRepository
}

func NewAddProject(consumer Consumer, repository domain.ProjectRepository) *AddProject {
	return &AddProject{consumer: consumer, repository: repository}
}

func (this AddProject) Run(i interface{}) interface{} {
	project := this.consumer.Consume(i).(domain.Project)
	log.Printf("AddProject.Run: %s", project.Name)
	this.repository.Add(project)
	return project
}
