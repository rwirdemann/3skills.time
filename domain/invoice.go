package domain

type Position struct {
	Hours float64
}

type Invoice struct {
	Positions map[int]map[string]Position
}

func NewInvoice() *Invoice {
	return &Invoice{make(map[int]map[string]Position)}
}

func NewInvoiceWithBookings(bookingsByProjectId map[int][]Booking) *Invoice {
	invoice := Invoice{make(map[int]map[string]Position)}
	for projectId, bookings := range bookingsByProjectId {
		if invoice.Positions[projectId] == nil {
			invoice.Positions[projectId] = make(map[string]Position)
		}
		for _, b := range bookings {
			position := invoice.Positions[projectId][b.Description]
			position.Hours += b.Hours
			invoice.Positions[projectId][b.Description] = position
		}
	}
	return &invoice
}

func (i *Invoice) AddPosition(projectId int, title string, hours float64) {
	if i.Positions[projectId] == nil {
		i.Positions[projectId] = make(map[string]Position)
	}
	i.Positions[projectId][title] = Position{Hours: hours}
}
