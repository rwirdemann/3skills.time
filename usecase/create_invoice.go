package usecase

import (
	"github.com/rwirdemann/3skills.time/foundation"
	"github.com/rwirdemann/3skills.time/domain"
)

type CreateInvoice struct {
	consumer   foundation.Consumer
	repository Repository
}

func (u CreateInvoice) Run(i ...interface{}) interface{} {
	return domain.Invoice{}
}