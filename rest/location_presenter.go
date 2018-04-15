package rest

import (
	"fmt"
)

func NewLocationPresenter() LocationPresenter {
	return LocationPresenter{}
}

type LocationPresenter struct {
}

func (p LocationPresenter) Present(i interface{}) interface{} {
	return fmt.Sprintf("Location: %s", i)
}
