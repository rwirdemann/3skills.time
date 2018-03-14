package rest

import (
	"encoding/json"
)

func NewJSONPresenter() JSONPresenter {
	return JSONPresenter{}
}

type JSONPresenter struct {
}

func (this JSONPresenter) Present(i interface{}) interface{} {
	b, _ := json.Marshal(i)
	return b
}
