package usecase

import (
	"github.com/rwirdemann/gotracker/foundation"
)

type GetBookings struct {
	consumer   foundation.Consumer
	presenter  foundation.Presenter
	repository Repository
}

func NewGetBookings(consumer foundation.Consumer,
	presenter foundation.Presenter,
	repository Repository) *GetBookings {
	return &GetBookings{consumer: consumer, presenter: presenter, repository: repository}
}

func (this GetBookings) Run(i interface{}) interface{} {
	projectId := this.consumer.Consume(i).(int)
	bookings := this.repository.AllBookings(projectId)
	return this.presenter.Present(bookings)
}
