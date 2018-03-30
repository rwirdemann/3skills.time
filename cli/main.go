package main

import (
	"fmt"

	"github.com/rwirdemann/3skills.time/cli/presenter"
	"github.com/rwirdemann/3skills.time/rest"

	"github.com/rwirdemann/3skills.time/database"
	"github.com/rwirdemann/3skills.time/usecase"
)

func main() {
	presenter := presenter.NewCLIPresenter()
	repository := database.NewMySQLRepository()

	consumer := rest.NewQueryConsumer("name")
	usecase := usecase.NewGetProjects(consumer, presenter, repository)
	result := usecase.Run(nil)
	fmt.Println(result)
}
