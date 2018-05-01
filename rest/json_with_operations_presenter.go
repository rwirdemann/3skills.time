package rest

import (
	"encoding/json"
	"github.com/rwirdemann/3skills.time/domain"
	"fmt"
)

func NewJSONProjectsWithOperationsPresenter() JSONProjectsWithOperations {
	return JSONProjectsWithOperations{}
}

type JSONProjectsWithOperations struct {
}

type Operation struct {
	Rel    string
	Method string
	Href   string
}

type ProjectWithOperations struct {
	domain.Project
	Operations []Operation
}

func (j JSONProjectsWithOperations) Present(i interface{}) interface{} {
	var withOperations []ProjectWithOperations
	projects := i.([]domain.Project)
	for _, p := range projects {
		show := Operation{"show", "GET", fmt.Sprintf("/projects/%d", p.Id)}
		update := Operation{"update", "PUT", fmt.Sprintf("/projects/%d", p.Id)}
		withOperation := ProjectWithOperations{Project: p, Operations: []Operation{show, update}}
		withOperations = append(withOperations, withOperation)
	}

	b, _ := json.Marshal(withOperations)
	return b
}
