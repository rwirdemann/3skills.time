package rest

import (
	"net/http"
)

type QueryConsumer struct {
	paramName string
}

func NewQueryConsumer(paramName string) QueryConsumer {
	return QueryConsumer{paramName: paramName}
}

func (q QueryConsumer) Consume(i interface{}) interface{} {
	request := i.(*http.Request)
	if filter, ok := request.URL.Query()[q.paramName]; ok {
		return filter[0]
	}
	return nil
}
