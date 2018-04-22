package usecase

import (
	"github.com/rwirdemann/3skills.time/foundation"
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

func (g GetProjects) Run(i ...interface{}) interface{} {
	var filter string
	switch v := g.consumer.Consume(i[0]).(type) {
	case string:
		filter = v
	}
	projects := g.repository.AllProjects(filter)
	return g.presenter.Present(projects)
}
