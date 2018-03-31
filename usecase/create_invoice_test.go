package usecase

import (
	"testing"
	"github.com/rwirdemann/3skills.time/domain"
	"fmt"
	"time"
)

func TestCreateInvoice(t *testing.T) {
	repository := NewFakeRepository()

	customerId := repository.AddCustomer(domain.Customer{Name: "3skills"})

	projectId := repository.AddProject(domain.Project{Name: "Energie", CustomerId: customerId})

	d := time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC)
	b1 := domain.Booking{Description: "Programmierung", ProjectId: projectId, Hours: 8, Date: d}
	b2 := domain.Booking{Description: "Programmierung", ProjectId: projectId, Hours: 8, Date: d}
	b3 := domain.Booking{Description: "Programmierung", ProjectId: projectId, Hours: 8, Date: d}
	b4 := domain.Booking{Description: "Programmierung", ProjectId: projectId, Hours: 8, Date: d}
	b5 := domain.Booking{Description: "Qualitätssicherung", ProjectId: projectId, Hours: 4, Date: d}

	repository.AddBooking(b1)
	repository.AddBooking(b2)
	repository.AddBooking(b3)
	repository.AddBooking(b4)
	repository.AddBooking(b5)

	projectId2 := repository.AddProject(domain.Project{Name: "Sicherheit", CustomerId: customerId})
	b6 := domain.Booking{Description: "Projektmanagement", ProjectId: projectId2, Hours: 8, Date: d}
	b7 := domain.Booking{Description: "Projektmanagement", ProjectId: projectId2, Hours: 8, Date: d}
	b8 := domain.Booking{Description: "Projektmanagement", ProjectId: projectId2, Hours: 8, Date: d}
	b9 := domain.Booking{Description: "Qualitätssicherung", ProjectId: projectId, Hours: 8, Date: d}

	repository.AddBooking(b6)
	repository.AddBooking(b7)
	repository.AddBooking(b8)
	repository.AddBooking(b9)

	usecase := NewCreateInvoice(NewIdentityConsumer(),
		NewIdentityConsumer(), NewIdentityConsumer(), NewIdentityPresenter(), repository)

	invoice := usecase.Run(customerId, 12, 2017)
	fmt.Printf("Invoice: %v", invoice)
}
