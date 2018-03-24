package usecase

import (
	"github.com/rwirdemann/gotracker-pg/foundation"
	"github.com/rwirdemann/gotracker/domain"
)

type AddBooking struct {
	projectIdConsumer foundation.Consumer
	bookingConsumer   foundation.Consumer
	repository        Repository
}

func NewAddBooking(projectIdConsumer foundation.Consumer,
	bookingConsumer foundation.Consumer,
	repository Repository) *AddBooking {
	return &AddBooking{projectIdConsumer: projectIdConsumer, bookingConsumer: bookingConsumer, repository: repository}
}

func (this AddBooking) Run(i ...interface{}) interface{} {
	projectId := this.projectIdConsumer.Consume(i[0]).(int)
	booking := this.bookingConsumer.Consume(i[1]).(*domain.Booking)
	booking.ProjectId = projectId
	this.repository.AddBooking(*booking)
	return booking
}
