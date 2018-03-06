package usecase

import "github.com/rwirdemann/go-tracker/design/domain"

type GetProjects struct {
	presenter  Presenter
	repository domain.ProjectRepository
}

func NewGetProjects(presenter Presenter, repository domain.ProjectRepository) *GetProjects {
	return &GetProjects{presenter: presenter, repository: repository}
}

func (this GetProjects) Run(i interface{}) interface{} {
	projects := this.repository.All()
	return this.presenter.Present(projects)
}
