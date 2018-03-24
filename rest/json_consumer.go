package rest

import (
	"encoding/json"
)

type JSONConsumer struct {
	result interface{}
}

func NewJSONConsumer(result interface{}) JSONConsumer {
	return JSONConsumer{result: result}
}

func (this JSONConsumer) Consume(body interface{}) interface{} {
	json.Unmarshal(body.([]byte), this.result)
	return this.result
}
