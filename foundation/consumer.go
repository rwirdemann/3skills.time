package foundation

type Consumer interface {
	Consume(i interface{}) interface{}
}
