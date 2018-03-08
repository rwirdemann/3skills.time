package main

import (
	"fmt"

	"github.com/rwirdemann/go-tracker/design/cli/presenter"

	"github.com/rwirdemann/go-tracker/design/database"
	"github.com/rwirdemann/go-tracker/design/usecase"
)

func main() {
	presenter := presenter.NewCLIPresenter()
	repository := database.NewMySQLProjectRepository()

	usecase := usecase.NewGetProjects(presenter, repository)
	result := usecase.Run(nil)
	fmt.Println(result)
}
