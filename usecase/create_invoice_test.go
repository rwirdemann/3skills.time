package usecase

import (
	"testing"
	"github.com/rwirdemann/3skills.time/domain"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	repository := NewFakeRepository()

	customerId := repository.AddCustomer(domain.Customer{Name: "3skills"})

	energie := repository.AddProject(domain.Project{Name: "Energie", CustomerId: customerId})
	b1 := booking(energie, "Programmierung", 8)
	b2 := booking(energie, "Programmierung", 8)
	b3 := booking(energie, "Programmierung", 8)
	b4 := booking(energie, "Programmierung", 8)
	b5 := booking(energie,"Qualitätssicherung", 3)
	repository.AddBookings(b1, b2, b3, b4, b5)

	sicherheit := repository.AddProject(domain.Project{Name: "Sicherheit", CustomerId: customerId})
	b6 := booking(sicherheit, "Projektmanagement",8)
	b7 := booking(sicherheit, "Projektmanagement",8)
	b8 := booking(sicherheit, "Projektmanagement",8)
	b9 := booking(sicherheit, "Qualitätssicherung",8)
	repository.AddBookings(b6, b7, b8, b9)

	usecase := NewCreateInvoice(NewIdentityConsumer(),
		NewIdentityConsumer(), NewIdentityConsumer(), NewIdentityPresenter(), repository)

	actual := usecase.Run(customerId, 12, 2017)

	expected := domain.NewInvoice()
	expected.AddPosition(energie, "Programmierung", 32)
	expected.AddPosition(energie, "Qualitätssicherung", 3)
	expected.AddPosition(sicherheit, "Projektmanagement", 24 )
	expected.AddPosition(sicherheit, "Qualitätssicherung", 8 )

	assert.Equal(t, expected, actual)
}

func booking(projectId int, description string, hours float64) domain.Booking {
	d := time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC)
	return domain.Booking{Description: description, ProjectId: projectId, Hours: hours, Date: d}
}
