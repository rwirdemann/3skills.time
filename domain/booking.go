package domain

import "time"

type Booking struct {
	Id          int
	ProjectId   int
	Description string
	Hours       float64
	Date		time.Time
}