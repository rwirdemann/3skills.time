package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rwirdemann/3skills.time/domain"
	"github.com/rwirdemann/3skills.time/usecase"
	"github.com/rwirdemann/3skills.time/database"
)

const emptyFilter = ""

func getProjects() ([]domain.Project, error) {
	consumer := usecase.IdentityConsumer{}
	presenter := usecase.IdentityPresenter{}
	repository := database.NewMySQLRepository()
	usecase := usecase.NewGetProjects(consumer, presenter, repository)

	return usecase.Run(emptyFilter).([]domain.Project), nil
}

func main() {
	lambda.Start(getProjects)
}
