package presenter

import (
	"fmt"

	"github.com/rwirdemann/gotracker/domain"
)

func NewCLIPresenter() CLIPresenter {
	return CLIPresenter{}
}

type CLIPresenter struct {
}

func (this CLIPresenter) Present(i interface{}) interface{} {
	result := "Projects:"
	projects := i.([]domain.Project)
	for i, p := range projects {
		result = fmt.Sprintf("%s\n%2d %s", result, i+1, p.Name)
	}
	return result
}
