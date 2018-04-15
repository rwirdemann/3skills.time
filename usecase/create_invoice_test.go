package usecase

import (
	"testing"
	"github.com/rwirdemann/3skills.time/domain"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/rwirdemann/3skills.time/foundation"
)

var identityConsumer foundation.Consumer
var identityPresenter foundation.Presenter

func init() {
	identityConsumer = NewIdentityConsumer()
	identityPresenter = NewIdentityPresenter()
}

func TestCreateInvoice(t *testing.T) {
	repository := NewFakeRepository()
	customerId := repository.AddCustomer(domain.Customer{Name: "3skills"})

	energy := domain.Project{Name: "Energie", CustomerId: customerId}
	energyId := repository.AddProject(energy)
	b1 := book(energyId, "Programmierung", 8)
	b2 := book(energyId, "Programmierung", 8)
	b3 := book(energyId, "Programmierung", 8)
	b4 := book(energyId, "Programmierung", 8)
	b5 := book(energyId, "Qualitätssicherung", 3)
	repository.AddBookings(b1, b2, b3, b4, b5)

	security := domain.Project{Name: "Sicherheit", CustomerId: customerId}
	securityId := repository.AddProject(security)
	b6 := book(securityId, "Projektmanagement", 8)
	b7 := book(securityId, "Projektmanagement", 8)
	b8 := book(securityId, "Projektmanagement", 8)
	b9 := book(securityId, "Qualitätssicherung", 8)
	repository.AddBookings(b6, b7, b8, b9)

	usecase := NewCreateInvoice(identityConsumer,
		identityConsumer, identityConsumer, identityPresenter, repository)
	actual := usecase.Run(customerId, 12, 2017)

	expected := domain.NewInvoice()
	expected.AddPosition(energyId, "Programmierung", 32)
	expected.AddPosition(energyId, "Qualitätssicherung", 3)
	expected.AddPosition(securityId, "Projektmanagement", 24)
	expected.AddPosition(securityId, "Qualitätssicherung", 8)

	assert.Equal(t, expected, actual)
}

func book(projectId int, description string, hours float64) domain.Booking {
	d := time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC)
	return domain.Booking{Description: description, ProjectId: projectId, Hours: hours, Date: d}
}
