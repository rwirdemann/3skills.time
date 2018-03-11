package rest

import (
	"net/http"
)

type QueryConsumer struct {
}

func NewQueryConsumer() QueryConsumer {
	return QueryConsumer{}
}

func (this QueryConsumer) Consume(i interface{}) interface{} {
	request := i.(*http.Request)
	if filter, ok := request.URL.Query()["name"]; ok {
		return filter[0]
	}
	return nil
}
