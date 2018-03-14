package main

import (
	"fmt"

	"github.com/rwirdemann/gotracker/cli/presenter"

	"github.com/rwirdemann/gotracker/database"
	"github.com/rwirdemann/gotracker/usecase"
)

func main() {
	presenter := presenter.NewCLIPresenter()
	repository := database.NewMySQLProjectRepository()

	usecase := usecase.NewGetProjects(presenter, repository)
	result := usecase.Run(nil)
	fmt.Println(result)
}
