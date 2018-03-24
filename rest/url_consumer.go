package rest

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// URLConsumer extracts the URL var "name" form the URL converts it's type to "typ"
type URLConsumer struct {
	name string
	typ  string
}

func NewURLConsumer(name string, typ string) URLConsumer {
	return URLConsumer{name: name, typ: typ}
}

func (this URLConsumer) Consume(i interface{}) interface{} {
	request := i.(*http.Request)
	if v, ok := mux.Vars(request)[this.name]; ok {
		switch this.typ {
		case "int":
			if i, err := strconv.Atoi(v); err == nil {
				return i
			}
		default:
			return i
		}
	}
	return nil
}
