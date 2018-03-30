package usecase

import (
	"github.com/rwirdemann/gotracker-pg/foundation"
	"github.com/rwirdemann/3skills.time/domain"
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

func (a AddBooking) Run(i ...interface{}) interface{} {
	projectId := a.projectIdConsumer.Consume(i[0]).(int)
	booking := a.bookingConsumer.Consume(i[1]).(*domain.Booking)
	booking.ProjectId = projectId
	a.repository.AddBooking(*booking)
	return booking
}
