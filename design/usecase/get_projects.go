package usecase

import (
	"github.com/rwirdemann/go-tracker/foundation"
)

type GetProjects struct {
	consumer   foundation.Consumer
	presenter  foundation.Presenter
	repository Repository
}

func NewGetProjects(consumer foundation.Consumer,
	presenter foundation.Presenter,
	repository Repository) *GetProjects {
	return &GetProjects{consumer: consumer, presenter: presenter, repository: repository}
}

func (this GetProjects) Run(i interface{}) interface{} {
	var filter string
	switch v := this.consumer.Consume(i).(type) {
	case string:
		filter = v
	}
	projects := this.repository.AllProjects(filter)
	return this.presenter.Present(projects)
}
