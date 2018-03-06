package usecase

type Consumer interface {
	Consume(i interface{}) interface{}
}
