package usecase

import (
	"github.com/rwirdemann/gotracker-pg/foundation"
	"github.com/rwirdemann/gotracker/domain"
)

type AddBooking struct {
	consumer   foundation.Consumer
	repository Repository
}

func NewAddBooking(consumer foundation.Consumer, repository Repository) *AddBooking {
	return &AddBooking{consumer: consumer, repository: repository}
}

func (this AddBooking) Run(i interface{}) interface{} {
	booking := this.consumer.Consume(i).(domain.Booking)
	this.repository.AddBooking(booking)
	return booking
}
