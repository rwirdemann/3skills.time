package usecase

import (
	"github.com/rwirdemann/3skills.time/foundation"
	"github.com/rwirdemann/3skills.time/domain"
)

type CreateInvoice struct {
	customerIdConsumer foundation.Consumer
	monthConsumer      foundation.Consumer
	yearConsumer       foundation.Consumer
	presenter          foundation.Presenter
	repository         Repository
}

func NewCreateInvoice(customerIdConsumer foundation.Consumer, monthConsumer foundation.Consumer,
	yearConsumer foundation.Consumer, presenter foundation.Presenter, repository Repository) *CreateInvoice {
	return &CreateInvoice{
		repository:         repository,
		presenter:          presenter,
		customerIdConsumer: customerIdConsumer,
		monthConsumer:      monthConsumer,
		yearConsumer:       yearConsumer}
}

func (u CreateInvoice) Run(i ...interface{}) interface{} {
	customerId := u.customerIdConsumer.Consume(i[0]).(int)
	month := u.monthConsumer.Consume(i[1]).(int)
	year := u.yearConsumer.Consume(i[2]).(int)

	if invoice, ok := u.repository.InvoiceByAndMonthAndYear(customerId, month, year); ok {
		return u.presenter.Present(invoice)
	}

	bookings := u.repository.AllBookingsByMonthAndYear(customerId, year, month)
	byProject := make(map[int][]domain.Booking)
	for _, b := range bookings {
		byProject[b.ProjectId] = append(byProject[b.ProjectId], b)
	}

	invoice := domain.NewInvoice(byProject)
	return u.presenter.Present(invoice)
}

