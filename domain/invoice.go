package domain

type Position struct {
	Title string
	Hours float64

}

type Invoice struct {
	positions map[int][]Position
}

func NewInvoice(bookingsByProjectId map[int][]Booking) *Invoice {
	invoice := Invoice{make(map[int][]Position)}
	for projectId, bookings := range bookingsByProjectId {
		for _, b := range bookings {
			p := Position{Title: b.Description, Hours: b.Hours}
			invoice.positions[projectId] = append(invoice.positions[projectId], p)
		}
	}
	return &invoice
}
